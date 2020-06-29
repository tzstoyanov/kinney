// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2020 VMware, Inc. Tzvetomir Stoyanov (VMware) <tz.stoyanov@gmail.com>

package api

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/CamusEnergy/kinney/controller/chargepoint/api/schema"
)

type ChargePointServer interface {
	GetLoad(*schema.GetLoadRequest) (*schema.GetLoadResponse, error)
	GetStations(*schema.GetStationsRequest) (*schema.GetStationsResponse, error)
	GetStationGroups(*schema.GetStationGroupsRequest) (*schema.GetStationGroupsResponse, error)
	GetCPNInstances(*schema.GetCPNInstancesRequest) (*schema.GetCPNInstancesResponse, error)
	ShedLoad(*schema.ShedLoadRequest) (*schema.ShedLoadResponse, error)
	ClearShedState(*schema.ClearShedStateRequest) (*schema.ClearShedStateResponse, error)
}

type apiRequest struct {
	payload interface{}
}

func (r *apiRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "getLoad":
		r.payload = &schema.GetLoadRequest{}
	case "getStations":
		r.payload = &schema.GetStationsRequest{}
	case "getStationGroups":
		r.payload = &schema.GetStationGroupsRequest{}
	case "getCPNInstances":
		r.payload = &schema.GetCPNInstancesRequest{}
	case "shedLoad":
		r.payload = &schema.ShedLoadRequest{}
	case "clearShedState":
		r.payload = &schema.ClearShedStateRequest{}
	default:
		return fmt.Errorf("unexpected request type: %#v", start.Name)
	}

	return d.DecodeElement(r.payload, &start)
}

func NewHandler(server ChargePointServer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		reqBytes, err := ioutil.ReadAll(req.Body)
		var request apiRequest

		if err := unmarshalEnvelope(reqBytes, nil, &request); err != nil {
			http.Error(w, "Failed unmarshaling the request", http.StatusBadRequest)
			return
		}

		var response interface{}
		switch req := request.payload.(type) {
		case *schema.GetLoadRequest:
			response, err = server.GetLoad(req)
		case *schema.GetStationsRequest:
			response, err = server.GetStations(req)
		case *schema.GetStationGroupsRequest:
			response, err = server.GetStationGroups(req)
		case *schema.GetCPNInstancesRequest:
			response, err = server.GetCPNInstances(req)
		case *schema.ShedLoadRequest:
			response, err = server.ShedLoad(req)
		case *schema.ClearShedStateRequest:
			response, err = server.ClearShedState(req)
		default:
			response, err = nil, fmt.Errorf("Not supported")
		}
		if err != nil {
			http.Error(w, "Failed handling the request", http.StatusMethodNotAllowed)
			return
		}

		if rBytes, err := marshalEnvelope(nil, response); err != nil {
			http.Error(w, "Failed marshaling the reply", http.StatusInternalServerError)
			return
		} else {
			w.Write(rBytes)
		}
	})
}
