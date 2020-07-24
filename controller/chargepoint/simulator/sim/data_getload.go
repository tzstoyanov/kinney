// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2020 VMware, Inc. Tzvetomir Stoyanov (VMware) <tz.stoyanov@gmail.com>

package sim

import (
	"fmt"

	"github.com/CamusEnergy/kinney/controller/chargepoint/api/schema"
)

func getStationLoad(group *chargeGroup, st *chargeStation, stId *string, res *schema.GetLoadResponse) (float32, error) {
	var sload float32
	sret := schema.GetLoadResponse_Station{StationID: *stId}
	if st.name != nil {
		sret.StationName = *st.name
	}
	if st.address != nil {
		sret.StationAddress = *st.address
	}

	simTime := group.getLoad.calcTime(group)
	for i, port := range st.ports {
		if p, v, err := group.getLoad.getPortLoad(port, simTime); err == nil {
			sload += p
			rport := schema.GetLoadResponse_Station_Port{
				PortNumber: i,
			}
			rport.PortLoadKW = fmt.Sprintf("%f", p)
			if v != nil {
				rport.UserID = *v.driverId
				rport.CredentialID = v.credentialID
			}
			sret.Ports = append(sret.Ports, rport)
		}
	}
	sret.StationLoadKW = fmt.Sprintf("%f", sload)
	res.Stations = append(res.Stations, sret)
	return sload, nil
}

func (e *EVChargers) getNextLoad(resp *schema.GetLoadResponse,
	group *chargeGroup, stationID *string) error {
	var gload float32

	if group == nil || stationID == nil {
		return fmt.Errorf("Station not found")
	}

	e.lock.Lock()
	defer e.lock.Unlock()

	if *stationID == "" {
		for i, s := range group.stations {
			if p, err := getStationLoad(group, s, &i, resp); err == nil {
				gload += p
			}
		}
	} else {
		if s, ok := group.stations[*stationID]; ok {
			if p, err := getStationLoad(group, s, stationID, resp); err == nil {
				gload += p
			}
		} else {
			return fmt.Errorf("Station not found")
		}
	}

	resp.StationGroupLoadKW = fmt.Sprintf("%f", gload)

	return nil
}
