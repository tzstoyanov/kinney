package schema

import "encoding/xml"

// API Guide (§ 8.5): "Use this call to retrieve organization and custom station
// that you have access rights to."
type GetOrgsAndStationGroupsRequest struct {
	XMLName xml.Name `xml:"urn:dictionary:com.chargepoint.webservices getOrgsAndStationGroups"`

	OrganizationID   string `xml:"searchQuery>orgID,omitempty"`
	OrganizationName string `xml:"searchQuery>organizationName,omitempty"`

	StationGroupID   int32 `xml:"searchQuery>sgID,omitempty"`
	StationGroupName int32 `xml:"searchQuery>sgName,omitempty"`
}

type GetOrgsAndStationGroupsResponse struct {
	XMLName xml.Name `xml:"urn:dictionary:com.chargepoint.webservices getOrgsAndStationGroupsResponse"`

	commonResponseParameters

	Organizations []GetOrgsAndStationGroupsResponse_Organization `xml:"orgData,omitempty"`
}

type GetOrgsAndStationGroupsResponse_Organization struct {
	OrganizationID   string `xml:"orgID,omitempty"`
	OrganizationName string `xml:"organizationName,omitempty"`

	StationGroups []GetOrgsAndStationGroupsResponse_Organization_StationGroup `xml:"sgData,omitempty"`
}

type GetOrgsAndStationGroupsResponse_Organization_StationGroup struct {
	StationGroupID   int32  `xml:"sgID,omitempty"`
	StationGroupName string `xml:"sgName,omitempty"`

	// API Guide (§ 8.5.3): "The Group ID of the parent group (`0` if it has
	// no parent group)."
	ParentGroupID string `xml:"parentGroupID,omitempty"`
}
