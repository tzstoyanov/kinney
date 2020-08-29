package schema

import "encoding/xml"

// API Guide (§ 10.1): "Use this method to get a list of connected or managed
// drivers.  If your organization uses ChargePoint Connections, this method will
// return the list of all drivers who have requested a connection with your
// organization, as well as a list of all drivers who have either been approved
// or rejected as connected drivers.  If your organization uses a branded
// ChargePoint portal to sign up drivers, then those drivers will appear in your
// management realm, and the list of those drivers will be returned by this
// method."
type GetUsersRequest struct {
	XMLName xml.Name `xml:"urn:dictionary:com.chargepoint.webservices getUsers"`

	// API Guide (§ 10.1.2): "Unique identifier of the driver.  This value
	// is not a driver account number or username."
	UserID string `xml:"searchQuery>userID,omitempty"`

	// API Guide (§ 10.1.2): "First name of the user/driver."
	FirstName string `xml:"searchQuery>firstName,omitempty"`
	// API Guide (§ 10.1.2): "Last name of the user/driver."
	LastName string `xml:"searchQuery>lastName,omitempty"`

	// API Guide (§ 10.1.2): "Find only records equal to or later than the
	// timestamp of the last change to any property of a user account.  This
	// property uses the ISO 8601 date time format in UTC
	// YYYY-MM-DDTHH:MM:SSZ."
	LastModified xsdDateTime `xml:"searchQuery>lastModifiedTimeStamp,omitempty"`

	// API Guide (§ 10.1.2): "Include this object in the query if you are
	// searching for a driver by their ChargePoint Connection properties.
	Connection *GetUsersRequest_Connection `xml:"searchQuery>Connection,omitempty"`

	// API Guide (§ 10.1.2): "Include this object if you wish to search for
	// users that are part of the Management Realm for your organization."
	ManagementRealm *GetUsersRequest_ManagementRealm `xml:"searchQuery>managementRealm,omitempty"`

	// API Guide (§ 10.1.2): "Use this property to search for users by
	// either the printed serial number from a ChargePoint RFID card or a
	// ChargePoint Mobile App identifier."
	CredentialID string `xml:"searchQuery>credentialID,omitempty"`

	// Undocumented pagination parameters.
	StartRecord int `xml:"searchQuery>startRecord,omitempty"`
	NumUsers    int `xml:"searchQuery>numUsers,omitempty"`
}

type UserConnectionStatus string

const (
	UserConnectionStatus_Approved     UserConnectionStatus = "APPROVED"
	UserConnectionStatus_NotConnected                      = "NOT CONNECTED"
	UserConnectionStatus_Rejected                          = "REJECTED"
	UserConnectionStatus_Pending                           = "PENDING"
)

type GetUsersRequest_Connection struct {
	// API Guide (§ 10.1.2): "Use this property to search for all users with
	// a given Connection status.  This property may be one of the following
	// values: APPROVED, NOT CONNECTED, REJECTED, PENDING."
	Status UserConnectionStatus `xml:"Status,omitempty"`

	// API Guide (§ 10.1.2): "This object contains a key-value-pair of
	// information that the driver provided when connecting to your
	// organization, such as an employee ID or club number."
	CustomInfo *GetUsersRequest_CustomInfo `xml:"customInfo,omitempty"`
}

type UserManagementRealmStatus string

const (
	UserManagementRealmStatus_Approved   UserManagementRealmStatus = "APPROVED"
	UserManagementRealmStatus_NotManaged                           = "NOT MANAGED"
)

type GetUsersRequest_ManagementRealm struct {
	// API Guide (§ 10.1.2): "Use this property to search for all users with
	// a given status.  This property may be one of the following values:
	// APPROVED, NOT MANAGED."
	Status UserManagementRealmStatus `xml:"Status,omitempty"`

	// API Guide (§ 10.1.2): "This object contains a key-value-pair of
	// information that the driver provided when joining the management
	// realm."
	CustomInfo *GetUsersRequest_CustomInfo `xml:"customInfo,omitempty"`
}

// This `xsd:complexType` is used for query both connection and management realm
// "custom info".  The documentation for each field is included separately
// below.
type GetUsersRequest_CustomInfo struct {
	// API Guide (§ 10.1.2): "The name of the key that you wish to use for
	// the search such as "Employee ID".  This key name is defined when your
	// Network Administrator defines a Connection Offer for your
	// organization."
	//
	// API Guide (§ 10.1.2): "The name of the key that you wish to use for
	// the search for users."
	Key string `xml:"Key,omitempty"`

	// API Guide (§ 10.1.2): "The value of this key that the user provided
	// when requesting a Connection with your organization."
	//
	// API Guide (§ 10.1.2): "The value of this key that the user provided
	// when signing up through your branded portal."
	Value string `xml:"Value,omitempty"`
}

type GetUsersResponse struct {
	XMLName xml.Name `xml:"urn:dictionary:com.chargepoint.webservices getUsersResponse"`

	commonResponseParameters

	Users    []GetUsersResponse_User `xml:"users>user,omitempty"`
	MoreFlag int                     `xml:"users>moreFlag,omitempty"`
}

type GetUsersResponse_User struct {
	// API Guide (§ 10.1.3): "Timestamp of the last change to any property
	// of this user account.  This property uses the ISO 8601 date time
	// format in UTC YYYY-MM-DDTHH:MM:SSZ."
	LastModified xsdDateTime `xml:"lastModifiedTimestamp,omitempty"`

	// API Guide (§ 10.1.3): "Unique identifier of the driver.  This value
	// is not a driver account number or username."
	UserID string `xml:"userID,omitempty"`

	// API Guide (§ 10.1.3): "First name of the user/driver."
	FirstName string `xml:"firstName,omitempty"`
	// API Guide (§ 10.1.3): "Last name of the user/driver."
	LastName string `xml:"lastName,omitempty"`

	// API Guide (§ 10.1.3): "If the user is Connected to or has requested a
	// Connection to your organization, this object will include properties
	// of that Connection."
	Connection *GetUsersResponse_Connection `xml:"Connection,omitempty"`

	// API Guide (§ 10.1.3): "If the user is part of your Management Realm,
	// this object will include the properties of that assiciation with your
	// organization."
	ManagementRealm *GetUsersResponse_ManagementRealm `xml:"managementRealm,omitempty"`

	// API Guide (§ 10.1.3): "The printed serial number from a ChargePoint
	// RFID card or a ChargePoint Mobile App identifier."
	CredentialIDs []string `xml:"credentialIDs>credentialID,omitempty"`

	RecordNumber int `xml:"recordNumber,omitempty"`
}

type GetUsersResponse_Connection struct {
	// API Guide (§ 10.1.3): "Status of the Connection between this user and
	// your organization.  This property will be one of the following
	// values: APPROVED, NOT CONNECTED, REJECTED, PENDING.
	Status UserConnectionStatus `xml:"Status,omitempty"`

	// API Guide (§ 10.1.3): "Time stamp indicated when the user requested a
	// connection with your organization."
	RequestTimestamp xsdDateTime `xml:"requestTimeStamp,omitempty"`

	// API Guide (§ 10.1.3): "This object contains a key-value-pair of
	// information that the dirver provided when connecting to your
	// organization, such as an employee ID or club number."
	CustomInfo []GetUsersResponse_CustomInfo `xml:"customInfos>customInfo,omitempty"`
}

type GetUsersResponse_ManagementRealm struct {
	// API Guide (§ 10.1.3): "Status of the user.  This property will be one
	// of the following values: APPROVED, NOT MANAGED"
	Status UserManagementRealmStatus `xml:"Status,omitempty"`

	// API Guide (§ 10.1.3): "Time stamp indicated when the user signed up
	// with your organization."
	SignupTimestamp xsdDateTime `xml:"signupTimeStamp,omitempty"`

	// API Guide (§ 10.1.3): "This object contains a key-value-pair of
	// information that the driver provided when signing up with your
	// organization."
	CustomInfo []GetUsersResponse_CustomInfo `xml:customInfos>customInfo,omitempty"`
}

// This `xsd:complexType` is used for query both connection and management realm
// "custom info".  The documentation for each field is included separately
// below.
type GetUsersResponse_CustomInfo struct {
	// API Guide (§ 10.1.3): "The name of the key for this custom property
	// of the connection.  This key name is defined when your Network
	// Administrator defines a Connection Offer for your organization."
	//
	// API Guide (§ 10.1.3): "The name of the custom key that you defined
	// for driver sign up for your Management Realm."
	Key string `xml:"Key,omitempty"`

	// API Guide (§ 10.1.3): "The value of this key that the user provided
	// when requesting a Connection with your organization."
	//
	// API Guide (§ 10.1.3): "The value of this key that the user provided
	// when signing up with your organization."
	Value string `xml:"Value,omitempty"`
}
