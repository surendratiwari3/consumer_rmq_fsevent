package model

/*
Conference Event
conference-end
conference-start

Participant Event
participant-leave
participant-join
participant-mute
participant-unmute
participant-hold
participant-unhold
participant-modify
participant-speech-start
participant-speech-stop

Announcement Event
announcement-end
announcement-fail*/
// ConferenceCall is the generic struct for handling common events.
type ConferenceEnd struct {
	ConferenceCommon
	CallSidEndingConference          string `json:"CallSidEndingConference,omitempty" form:"CallSidEndingConference" query:"CallSidEndingConference"`
	ParticipantLabelEndingConference string `json:"ParticipantLabelEndingConference,omitempty" form:"ParticipantLabelEndingConference" query:"ParticipantLabelEndingConference"`
	ReasonConferenceEnded            string `json:"ReasonConferenceEnded,omitempty" form:"ReasonConferenceEnded" query:"ReasonConferenceEnded"`
	Reason                           string `json:"Reason,omitempty" form:"Reason" query:"Reason"`
}

type ConferenceCommon struct {
	ConferenceSid       string `json:"ConferenceSid,omitempty" form:"ConferenceSid" query:"ConferenceSid"`
	FriendlyName        string `json:"FriendlyName,omitempty" form:"FriendlyName" query:"FriendlyName"`
	AccountSid          string `json:"AccountSid,omitempty" form:"AccountSid" query:"AccountSid"`
	Timestamp           string `json:"Timestamp,omitempty" form:"Timestamp" query:"Timestamp"`
	SequenceNumber      string `json:"SequenceNumber,omitempty" form:"SequenceNumber" query:"SequenceNumber"`
	StatusCallbackEvent string `json:"StatusCallbackEvent,omitempty" form:"StatusCallbackEvent" query:"StatusCallbackEvent"`
}

type ConferenceParticipant struct {
	ConferenceCommon
	CallSid                string `json:"CallSid,omitempty" form:"CallSid" query:"CallSid"`
	Muted                  string `json:"Muted,omitempty" form:"Muted" query:"Muted"`
	Hold                   string `json:"Hold,omitempty" form:"Hold" query:"Hold"`
	Coaching               string `json:"Coaching,omitempty" form:"Coaching" query:"Coaching"`
	EndConferenceOnExit    string `json:"EndConferenceOnExit,omitempty" form:"EndConferenceOnExit" query:"EndConferenceOnExit"`
	StartConferenceOnEnter string `json:"StartConferenceOnEnter,omitempty" form:"StartConferenceOnEnter" query:"StartConferenceOnEnter"`
}

/*
	Announcement using api for participant or announcement to conference
*/
type ConferenceAnnouncement struct {
	ConferenceCommon
	AnnounceUrl              string `json:"AnnounceUrl,omitempty" form:"AnnounceUrl" query:"AnnounceUrl"`
	ReasonAnnouncementFailed string `json:"ReasonAnnouncementFailed,omitempty" form:"ReasonAnnouncementFailed" query:"ReasonAnnouncementFailed"`
}
