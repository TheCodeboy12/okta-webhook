package structs

import (
	"fmt"
	"time"
)

type OktaEventHook struct { // Changed struct name to match file name
	EventType          string    `json:"eventType,omitempty"`
	EventTypeVersion   string    `json:"eventTypeVersion,omitempty"`
	CloudEventsVersion string    `json:"cloudEventsVersion,omitempty"`
	Source             string    `json:"source,omitempty"`
	EventID            string    `json:"eventId,omitempty"`
	Data               Data      `json:"data,omitempty"`
	EventTime          time.Time `json:"eventTime,omitempty"`
	ContentType        string    `json:"contentType,omitempty"`
}

func (e *OktaEventHook) Validate() error {

	if len(e.Data.Events) == 0 {
		return fmt.Errorf("event has no targets")
	}
	return nil
}

type Data struct {
	Events []Event `json:"events,omitempty"`
}

type Event struct {
	UUID                  string                `json:"uuid,omitempty"`
	Published             time.Time             `json:"published"`
	EventType             string                `json:"eventType"`
	Version               string                `json:"version"`
	DisplayMessage        string                `json:"displayMessage"`
	Severity              string                `json:"severity"`
	Client                Client                `json:"client"`
	Device                interface{}           `json:"device"` // Can be null
	Actor                 Actor                 `json:"actor"`
	Outcome               Outcome               `json:"outcome"`
	Target                []Target              `json:"target"`
	Transaction           Transaction           `json:"transaction"`
	DebugContext          DebugContext          `json:"debugContext"`
	LegacyEventType       string                `json:"legacyEventType"`
	AuthenticationContext AuthenticationContext `json:"authenticationContext"`
	SecurityContext       SecurityContext       `json:"securityContext"`
	InsertionTimestamp    interface{}           `json:"insertionTimestamp"` // Can be null
}

type UserAgent struct {
	RawUserAgent string `json:"rawUserAgent,omitempty"`
	OS           string `json:"os,omitempty"`
	Browser      string `json:"browser,omitempty"`
}

type GeographicalContext struct {
	City        string       `json:"city"`
	State       string       `json:"state"`
	Country     string       `json:"country"`
	PostalCode  interface{}  `json:"postalCode"`            // Can be null
	Geolocation *Geolocation `json:"geolocation,omitempty"` // Added omitempty
}

type Geolocation struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type IPChain struct {
	IP                  string               `json:"ip"`
	GeographicalContext *GeographicalContext `json:"geographicalContext,omitempty"`
	Version             string               `json:"version"`
	Source              interface{}          `json:"source,omitempty"` // Can be null, added omitempty
}

type Client struct {
	UserAgent           UserAgent            `json:"userAgent"`
	Zone                string               `json:"zone"`
	Device              string               `json:"device"`
	ID                  interface{}          `json:"id"` // Can be null
	IPAddress           string               `json:"ipAddress,omitempty"`
	GeographicalContext *GeographicalContext `json:"geographicalContext,omitempty"`
	IPChain             []IPChain            `json:"ipChain,omitempty"`
}

type Actor struct {
	ID          string      `json:"id,omitempty"`
	Type        string      `json:"type,omitempty"`
	AlternateID string      `json:"alternateId,omitempty"`
	DisplayName string      `json:"displayName,omitempty"`
	DetailEntry interface{} `json:"detailEntry,omitempty"` // Can be null, added omitempty
}

type Outcome struct {
	Result string      `json:"result,omitempty"`
	Reason interface{} `json:"reason,omitempty"` // Can be null, added omitempty
}

type Target struct {
	ID          string      `json:"id,omitempty"`
	Type        string      `json:"type,omitempty"`
	AlternateID string      `json:"alternateId,omitempty"`
	DisplayName string      `json:"displayName,omitempty"`
	DetailEntry interface{} `json:"detailEntry,omitempty"` // Can be null, added omitempty
}

type Transaction struct {
	Type   string                 `json:"type,omitempty"`
	ID     string                 `json:"id,omitempty"`
	Detail map[string]interface{} `json:"detail,omitempty"`
}

type DebugData struct {
	Appname    string `json:"appname,omitempty"`
	RequestID  string `json:"requestId,omitempty"`
	DtHash     string `json:"dtHash,omitempty"`
	RequestUri string `json:"requestUri,omitempty"`
	URL        string `json:"url,omitempty"`
}

type DebugContext struct {
	DebugData DebugData `json:"debugData,omitempty"`
}

type AuthenticationContext struct {
	AuthenticationProvider interface{} `json:"authenticationProvider"` // Can be null
	CredentialProvider     interface{} `json:"credentialProvider"`     // Can be null
	CredentialType         interface{} `json:"credentialType"`         // Can be null
	Issuer                 interface{} `json:"issuer"`                 // Can be null
	AuthenticationStep     int         `json:"authenticationStep,omitempty"`
	RootSessionId          string      `json:"rootSessionId,omitempty"`
	ExternalSessionId      string      `json:"externalSessionId,omitempty"`
	AuthenticatorContext   interface{} `json:"authenticatorContext"` // Can be null
	Interface              interface{} `json:"interface"`            // Can be null
}

type SecurityContext struct {
	AsNumber int         `json:"asNumber,omitempty"`
	AsOrg    interface{} `json:"asOrg"` // Can be null
	Isp      string      `json:"isp,omitempty"`
	Domain   string      `json:"domain,omitempty"`
	IsProxy  bool        `json:"isProxy,omitempty"`
}
