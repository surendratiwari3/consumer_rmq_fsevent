package model

type ConferenceFsEvent struct {
	Action                         string `json:"Action"`
	AnswerState                    string `json:"Answer-State"`
	CallDirection                  string `json:"Call-Direction"`
	CallerAni                      string `json:"Caller-Ani"`
	CallerCalleeIDName             string `json:"Caller-Callee-Id-Name"`
	CallerCalleeIDNumber           string `json:"Caller-Callee-Id-Number"`
	CallerCallerIDName             string `json:"Caller-Caller-Id-Name"`
	CallerCallerIDNumber           string `json:"Caller-Caller-Id-Number"`
	CallerChannelAnsweredTime      string `json:"Caller-Channel-Answered-Time"`
	CallerChannelBridgedTime       string `json:"Caller-Channel-Bridged-Time"`
	CallerChannelCreatedTime       string `json:"Caller-Channel-Created-Time"`
	CallerChannelHangupTime        string `json:"Caller-Channel-Hangup-Time"`
	CallerChannelHoldAccum         string `json:"Caller-Channel-Hold-Accum"`
	CallerChannelLastHold          string `json:"Caller-Channel-Last-Hold"`
	CallerChannelName              string `json:"Caller-Channel-Name"`
	CallerChannelProgressMediaTime string `json:"Caller-Channel-Progress-Media-Time"`
	CallerChannelProgressTime      string `json:"Caller-Channel-Progress-Time"`
	CallerChannelResurrectTime     string `json:"Caller-Channel-Resurrect-Time"`
	CallerChannelTransferTime      string `json:"Caller-Channel-Transfer-Time"`
	CallerContext                  string `json:"Caller-Context"`
	CallerDestinationNumber        string `json:"Caller-Destination-Number"`
	CallerDirection                string `json:"Caller-Direction"`
	CallerLogicalDirection         string `json:"Caller-Logical-Direction"`
	CallerNetworkAddr              string `json:"Caller-Network-Addr"`
	CallerOrigCallerIDName         string `json:"Caller-Orig-Caller-Id-Name"`
	CallerOrigCallerIDNumber       string `json:"Caller-Orig-Caller-Id-Number"`
	CallerPrivacyHideName          string `json:"Caller-Privacy-Hide-Name"`
	CallerPrivacyHideNumber        string `json:"Caller-Privacy-Hide-Number"`
	CallerProfileCreatedTime       string `json:"Caller-Profile-Created-Time"`
	CallerProfileIndex             string `json:"Caller-Profile-Index"`
	CallerScreenBit                string `json:"Caller-Screen-Bit"`
	CallerSource                   string `json:"Caller-Source"`
	CallerUniqueID                 string `json:"Caller-Unique-Id"`
	ChannelCallState               string `json:"Channel-Call-State"`
	ChannelCallUUID                string `json:"Channel-Call-Uuid"`
	ChannelHitDialplan             string `json:"Channel-Hit-Dialplan"`
	ChannelName                    string `json:"Channel-Name"`
	ChannelReadCodecBitRate        string `json:"Channel-Read-Codec-Bit-Rate"`
	ChannelReadCodecName           string `json:"Channel-Read-Codec-Name"`
	ChannelReadCodecRate           string `json:"Channel-Read-Codec-Rate"`
	ChannelState                   string `json:"Channel-State"`
	ChannelStateNumber             string `json:"Channel-State-Number"`
	ChannelWriteCodecBitRate       string `json:"Channel-Write-Codec-Bit-Rate"`
	ChannelWriteCodecName          string `json:"Channel-Write-Codec-Name"`
	ChannelWriteCodecRate          string `json:"Channel-Write-Codec-Rate"`
	ConferenceDomain               string `json:"Conference-Domain"`
	ConferenceGhosts               string `json:"Conference-Ghosts"`
	ConferenceName                 string `json:"Conference-Name"`
	ConferenceProfileName          string `json:"Conference-Profile-Name"`
	ConferenceSize                 string `json:"Conference-Size"`
	ConferenceUniqueID             string `json:"Conference-Unique-Id"`
	CoreUUID                       string `json:"Core-Uuid"`
	CurrentEnergy                  string `json:"Current-Energy"`
	EnergyLevel                    string `json:"Energy-Level"`
	EventCallingFile               string `json:"Event-Calling-File"`
	EventCallingFunction           string `json:"Event-Calling-Function"`
	EventCallingLineNumber         string `json:"Event-Calling-Line-Number"`
	EventDateGmt                   string `json:"Event-Date-Gmt"`
	EventDateLocal                 string `json:"Event-Date-Local"`
	EventDateTimestamp             string `json:"Event-Date-Timestamp"`
	EventName                      string `json:"Event-Name"`
	EventSequence                  string `json:"Event-Sequence"`
	EventSubclass                  string `json:"Event-Subclass"`
	Floor                          string `json:"Floor"`
	FreeswitchHostname             string `json:"Freeswitch-Hostname"`
	FreeswitchIpv4                 string `json:"Freeswitch-Ipv4"`
	FreeswitchIpv6                 string `json:"Freeswitch-Ipv6"`
	FreeswitchSwitchname           string `json:"Freeswitch-Switchname"`
	Hear                           string `json:"Hear"`
	Hold                           string `json:"Hold"`
	MemberGhost                    string `json:"Member-Ghost"`
	MemberID                       string `json:"Member-Id"`
	MemberType                     string `json:"Member-Type"`
	MuteDetect                     string `json:"Mute-Detect"`
	PresenceCallDirection          string `json:"Presence-Call-Direction"`
	See                            string `json:"See"`
	Speak                          string `json:"Speak"`
	Talking                        string `json:"Talking"`
	UniqueID                       string `json:"Unique-Id"`
	Video                          string `json:"Video"`
	StatuscallbackMethod           string `json:"Statuscallback_Method"`
	StatuscallbackURL              string `json:"Statuscallback_Url"`
	AuthSid						   string `json:"AuthSid"`
}
