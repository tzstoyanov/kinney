// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2020 VMware, Inc. Tzvetomir Stoyanov (VMware) <tz.stoyanov@gmail.com>

package sim

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type StationPort struct {
	PortNumber   string  `json:"portNumber"`
	UserID       string  `json:"userID"`
	CredentialID string  `json:"credentialID"`
	ShedState    int     `json:"shedState"`
	PortLoad     float32 `json:"portLoad"`
	AllowedLoad  float32 `json:"allowedLoad"`
	PercentShed  int     `json:"percentShed"`
}

type StationData struct {
	StationID   string        `json:"stationID"`
	StationName string        `json:"stationName"`
	Address     string        `json:"Address"`
	StationLoad float32       `json:"stationLoad"`
	Ports       []StationPort `json:"Port"`
}

type ChargeData struct {
	ResponseCode string        `json:"responseCode"`
	ResponseText string        `json:"responseText"`
	SgID         int           `json:"sgID"`
	NumStations  int           `json:"numStations"`
	GroupName    string        `json:"groupName"`
	SgLoad       string        `json:"sgLoad"`
	Stations     []StationData `json:"stationData"`
}

type ChargeRecord struct {
	Ts   float64    `json:"ts"`
	Data ChargeData `json:"data"`
}

type records struct {
	Records []ChargeRecord `json:"records"`
}

func json2ev(jrec *records, e *EVChargers) int {
	var samples int

	cpnID := "1"
	cpnName := cpnDefName
	cpnDesc := cpnDefDesc
	e.getCPNetwork(&cpnID, &cpnName, &cpnDesc)

	vmGroup := vmwareOrganizationID
	vmName := vmwareOrganizationName
	e.getChargeFacility(&vmGroup, &vmName)

	for _, r := range jrec.Records {
		if r.Data.ResponseCode != "100" {
			continue
		}
		sgid := strconv.Itoa(r.Data.SgID)
		e.getChargeGroup(&vmGroup, &sgid, &r.Data.GroupName, &getLoadReplay{})

		for _, s := range r.Data.Stations {
			e.getChargeStation(&vmGroup, &sgid, &s.StationID, &s.StationName,
				&s.Address, &locGeo{lat: geoDefLat, long: geoDefLong})
			for _, p := range s.Ports {
				e.getChargePort(&vmGroup, &sgid, &s.StationID, &p.PortNumber)
				if p.PortLoad != 0 && p.CredentialID != "" {
					sec, dec := math.Modf(r.Ts)
					time := time.Unix(int64(sec), int64(dec*(1e9)))
					e.addChargeProbe(&vmGroup, &sgid, &s.StationID,
						&p.PortNumber, &p.UserID, &p.CredentialID, time, p.PortLoad)
					samples++
				}
			}
		}
	}

	return samples
}

func jsonLoad(fname *string, e *EVChargers) (int, error) {
	var recs records

	if file, err := os.Open(*fname); err != nil {
		return 0, err
	} else {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err != nil {
			return 0, err
		} else if err := json.Unmarshal(data, &recs); err != nil {
			return 0, err
		}
	}

	return json2ev(&recs, e), nil
}

func DataLoadJsonDir(dir *string, e *EVChargers) (int, error) {
	var samples int

	err := filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || filepath.Ext(strings.TrimSpace(path)) != ".json" {
			return nil
		}

		if numSamples, err := jsonLoad(&path, e); err == nil {
			samples += numSamples
		} else {
			log.Println("Failed to load ", path, ":", err)
		}

		return nil
	})

	sortReplaySamples(e)

	return samples, err
}
