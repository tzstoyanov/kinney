package schema

import "encoding/xml"

// API Guide (§ 4.1): "Use this call to retrieve ChargePoint NOS instances."
//
// TODO(james): Figure out what "NOS" means, as it is not documented in the API
// Guide.
type GetCPNInstancesRequest struct {
	XMLName xml.Name `xml:"urn:dictionary:com.chargepoint.webservices getCPNInstances"`
}

type GetCPNInstancesResponse struct {
	XMLName xml.Name `xml:"urn:dictionary:com.chargepoint.webservices getCPNInstancesResponse"`

	ChargePointNetworks []GetCPNInstancesResponse_ChargePointNetwork `xml:"CPN,omitempty"`
}

type GetCPNInstancesResponse_ChargePointNetwork struct {
	ID          string `xml:"cpnID,omitempty"`
	Name        string `xml:"cpnName,omitempty"`
	Description string `xml:"cpnDescription,omitempty"`
}
