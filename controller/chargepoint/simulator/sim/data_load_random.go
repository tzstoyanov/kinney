// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2020 VMware, Inc. Tzvetomir Stoyanov (VMware) <tz.stoyanov@gmail.com>

package sim

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sethvargo/go-diceware/diceware"
)

type fixedStation struct {
	Name    string    `json:"Name,omitempty"`
	Address string    `json:"Address,omitempty"`
	Geo     []float32 `json:"Geo,omitempty"`
	Ports   []float32 `json:"Ports,omitempty"`
}

type fixedGroup struct {
	Name     string         `json:"Name,omitempty"`
	Stations []fixedStation `json:"Stations,omitempty"`
}

type fixedFacility struct {
	Name   string       `json:"Name,omitempty"`
	Groups []fixedGroup `json:"Groups,omitempty"`
}

type fixedCPN struct {
	Name string `json:"Name,omitempty"`
	Desc string `json:"Description,omitempty"`
}

type randomParams struct {
	MaxCPNs           int             `json:"maxCPNs,omitempty"`
	MaxFacilities     int             `json:"maxFacilities,omitempty"`
	MaxChargeGroups   int             `json:"maxChargeGroups,omitempty"`
	MaxChargeStations int             `json:"maxChargeStations,omitempty"`
	MaxChargePorts    int             `json:"maxChargePorts,omitempty"`
	MaxVehicleBattery int             `json:"maxVehicleBattery,omitempty"`
	NumCPNs           int             `json:"CPNs,omitempty"`
	NumFacilities     int             `json:"Facilities,omitempty"`
	NumChargeGroups   int             `json:"ChargeGroups,omitempty"`
	NumChargeStations int             `json:"ChargeStations,omitempty"`
	NumChargePorts    int             `json:"ChargePorts,omitempty"`
	PortLoad          int             `json:"PortLoad,omitempty"`
	VehicleUnplug     int             `json:"VehicleUnplug,omitempty"`
	RandSeed          int             `json:"RandomSeed,omitempty"`
	CPNs              []fixedCPN      `json:"FixedCPN,omitempty"`
	Facilities        []fixedFacility `json:"FixedFacilities,omitempty"`
}

const (
	defCPNs           = 1
	defFacilities     = 4
	defChargeGroups   = 4
	defChargeStations = 8
	defChargePorts    = 2
	defVehicleBattery = 80  // KW
	defPortCapacity   = 8.0 // KW

	// default probability of a port to have a charging session, in %
	defPortLoad = 50

	// default probability of a vehicle to unplug before fully charged, in %
	defVehicleUnplug = 20
)

func randParam(min int, max int) int {
	if max <= min {
		return min
	}
	return min + rand.Intn(max-min)
}

func randString(words int, sep string) *string {
	l, err := diceware.Generate(words)
	str := ""
	if err != nil {
		return &str
	}
	str = strings.Join(l, sep)
	return &str
}

func newCPN(e *EVChargers, c int, name, desc *string) {
	str := strconv.FormatInt(int64(c), 10)
	e.getCPNetwork(&str, name, desc)
}

func newPort(e *EVChargers, facility, group, station *string, id int, capacity float32) {
	str := strconv.FormatInt(int64(id), 10)
	port := e.getChargePort(facility, group, station, &str)
	port.max_capacity = capacity
	port.current_capacity = capacity
}

func newStation(e *EVChargers, facility, group *string, rParam *randomParams, fStation *fixedStation) {
	nChargePorts := rParam.NumChargePorts
	if rParam.MaxChargePorts > 0 {
		nChargePorts = randParam(1, rParam.MaxChargeStations)
	}

	var sAddress, sName string
	var loc locGeo
	sID := fmt.Sprintf("1:%d", randParam(100000, 999999))
	if fStation != nil {
		sName = fStation.Name
		sAddress = fStation.Address
		if len(fStation.Geo) == 2 {
			loc = locGeo{}
			loc.lat = fmt.Sprintf("%f", fStation.Geo[0])
			loc.long = fmt.Sprintf("%f", fStation.Geo[1])
		} else {
			loc = locGeo{lat: geoDefLat, long: geoDefLong}
		}
	} else {
		sName = *randString(1, "")
		sAddress = *randString(5, ",")
		loc = locGeo{lat: geoDefLat, long: geoDefLong}
	}

	e.getChargeStation(facility, group, &sID, &sName, &sAddress, &loc)
	if fStation != nil && len(fStation.Ports) > 0 {
		for i, p := range fStation.Ports {
			newPort(e, facility, group, &sID, i, p)
		}
	} else {
		for p := 0; p < nChargePorts; p++ {
			newPort(e, facility, group, &sID, p, defPortCapacity)
		}
	}
}

func newGroup(e *EVChargers, facility *string, rParam *randomParams, fGroup *fixedGroup) {
	nChargeStations := rParam.NumChargeStations
	if rParam.MaxChargeStations > 0 {
		nChargeStations = randParam(1, rParam.MaxChargeStations)
	}

	var gName string
	gID := fmt.Sprintf("%d", randParam(100000, 999999))
	if fGroup != nil {
		gName = fGroup.Name
	} else {
		gName = *randString(1, "")
	}
	e.getChargeGroup(facility, &gID, &gName,
		&getLoadRandom{
			portLoad:          rParam.PortLoad,
			vehicleUnplug:     rParam.VehicleUnplug,
			maxVehicleBattery: rParam.MaxVehicleBattery,
		})
	if fGroup != nil && len(fGroup.Stations) > 0 {
		for _, s := range fGroup.Stations {
			newStation(e, facility, &gID, rParam, &s)
		}
	} else {
		for s := 0; s < nChargeStations; s++ {
			newStation(e, facility, &gID, rParam, nil)
		}
	}
}

func newFacility(e *EVChargers, rParam *randomParams, fFacility *fixedFacility) {
	nChargeGroups := rParam.NumChargeGroups
	if rParam.MaxChargeGroups > 0 {
		nChargeGroups = randParam(1, rParam.MaxChargeGroups)
	}

	var fName string
	fID := fmt.Sprintf("1:%d", randParam(10000000, 99999999))
	if fFacility != nil {
		fName = fFacility.Name
	} else {
		fName = *randString(1, "")
	}
	e.getChargeFacility(&fID, &fName)
	if fFacility != nil && len(fFacility.Groups) > 0 {
		for _, g := range fFacility.Groups {
			newGroup(e, &fID, rParam, &g)
		}
	} else {
		for g := 0; g < nChargeGroups; g++ {
			newGroup(e, &fID, rParam, nil)
		}
	}
}

func genRandom(param *randomParams, e *EVChargers) {

	nCPNs := param.NumCPNs
	nFacilities := param.NumFacilities

	if param.RandSeed > 0 {
		rand.Seed(int64(param.RandSeed))
	} else {
		rand.Seed(int64(time.Now().Nanosecond()))
	}

	if param.MaxCPNs > 0 {
		nCPNs = randParam(1, param.MaxCPNs)
	}
	if param.MaxFacilities > 0 {
		nFacilities = randParam(1, param.MaxFacilities)
	}

	// Generate networks
	if len(param.CPNs) > 0 {
		for c, n := range param.CPNs {
			newCPN(e, c, &n.Name, &n.Desc)
		}
	} else {
		for c := 0; c <= nCPNs; c++ {
			newCPN(e, c, randString(1, ""), randString(5, ","))
		}
	}

	// Generate facilities
	if len(param.Facilities) > 0 {
		for _, f := range param.Facilities {
			newFacility(e, param, &f)
		}
	} else {
		for f := 0; f < nFacilities; f++ {
			newFacility(e, param, nil)
		}
	}
}

func DataLoadRandom(config *string, e *EVChargers) error {
	var err error
	var file *os.File
	var data []byte
	params := randomParams{
		NumCPNs:           defCPNs,
		NumFacilities:     defFacilities,
		NumChargeGroups:   defChargeGroups,
		NumChargeStations: defChargeStations,
		NumChargePorts:    defChargePorts,
		MaxVehicleBattery: defVehicleBattery,
		PortLoad:          defPortLoad,
		VehicleUnplug:     defVehicleUnplug,
	}

	if file, err = os.Open(*config); err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	if data, err = ioutil.ReadAll(file); err != nil {
		fmt.Println(err)
		return err
	}

	if err = json.Unmarshal(data, &params); err != nil {
		fmt.Println(err)
		return err
	}

	genRandom(&params, e)
	return nil
}
