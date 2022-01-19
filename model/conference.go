package model

type ConferenceStatus string

// ConferenceCall is the generic struct for handling common events.
type ConferenceCall struct {
	ConferenceSid       string `json:"ConferenceSid"`
	FriendlyName        string
	AccountSid          string `json:"AccountSid"`
	SequenceNumber      uint
	StatusCallbackEvent ConferenceStatus
	CallSID             *string `json:"CallSid,omitempty"`
}