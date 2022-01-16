package model

type ConferenceStatus string

/*
	Read from redis for conference attribute
*/
type ConferenceDetailsFromCache struct {
	DialConferenceStatusCallback       string `json:"DialConferenceStatusCallback"`
	DialConferenceStatusCallbackMethod string `json:"DialConferenceStatusCallbackMethod"`
	DialConfSid                        string `json:"DialConfSid"`
	DialConfAccountSid                 string `json:"DialConfAccountSid"`
}

// ConferenceCall is the generic struct for handling common events.
type ConferenceCall struct {
	ConferenceSid       string `json:"ConferenceSid"`
	FriendlyName        string
	AccountSid          string `json:"AccountSid"`
	SequenceNumber      uint
	StatusCallbackEvent ConferenceStatus
	CallSID             *string `json:"CallSid,omitempty"`
}
