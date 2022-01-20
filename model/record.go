package model

type RecordCallback struct {
	RecordEventTimestamp string `json:"Timestamp,omitempty" form:"Timestamp" query:"Timestamp" freeswitch:"Event-Date-GMT"`
	RecordingSource      string `json:"RecordingSource,omitempty" form:"RecordingSource" query:"RecordingSource"`
	RecordingTrack       string `json:"RecordingTrack,omitempty" form:"RecordingTrack" query:"RecordingTrack"`
	RecordingSid         string `json:"RecordingSid,omitempty" form:"RecordingSid" query:"RecordingSid"`
	RecordingUrl         string `json:"RecordingUrl,omitempty" form:"RecordingUrl" query:"RecordingUrl" freeswitch:"Record-File-Path"`
	RecordingStatus      string `json:"RecordingStatus,omitempty" form:"RecordingStatus" query:"RecordingStatus"`
	RecordingChannels    string `json:"RecordingChannels,omitempty" form:"RecordingChannels" query:"RecordingChannels"`
	ErrorCode            string `json:"ErrorCode,omitempty" form:"ErrorCode" query:"ErrorCode"`
	RecordCallSid        string `json:"CallSid,omitempty" form:"CallSid" query:"CallSid"`
	RecordingStartTime   string `json:"RecordingStartTime,omitempty" form:"RecordingStartTime" query:"RecordingStartTime"`
	RecordAccountSid     string `json:"AccountSid,omitempty" form:"AccountSid" query:"AccountSid"`
	RecordingDuration    string `json:"RecordingDuration,omitempty" form:"RecordingDuration" query:"RecordingDuration" freeswitch:"Variable_record_seconds"`
}
