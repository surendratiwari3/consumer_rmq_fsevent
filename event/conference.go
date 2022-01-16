package event

import (
	"encoding/json"
	"fmt"
	"github.com/consumer_rmq_fsevent/httprest"
	"github.com/consumer_rmq_fsevent/model"
	"github.com/consumer_rmq_fsevent/redis"
	"log"
	"net/url"
)

type ConfEventInterface interface {
	ProcessConfEvent(eventData []byte) error
}

type ConfEventHandler struct {
	httpHandler  httprest.RestInterface
	cacheHandler redis.CacheInterface
}

type ConferenceStatus string

// ConferenceStatus Enumerations
const (
	StatusConferenceEnd          ConferenceStatus = "conference-end"
	StatusConferenceStart        ConferenceStatus = "conference-start"
	StatusParticipantLeave       ConferenceStatus = "participant-leave"
	StatusParticipantJoin        ConferenceStatus = "participant-join"
	StatusParticipantMute        ConferenceStatus = "participant-mute"
	StatusParticipantUnmute      ConferenceStatus = "participant-unmute"
	StatusParticipantHold        ConferenceStatus = "participant-hold"
	StatusParticipantUnhold      ConferenceStatus = "participant-unhold"
	StatusParticipantSpeechStart ConferenceStatus = "participant-speech-start"
	StatusParticipantSpeechStop  ConferenceStatus = "participant-speech-stop"
	StatusAnnouncementEnd        ConferenceStatus = "announcement-end"
	StatusAnnouncementFail       ConferenceStatus = "announcement-fail"
)

func NewConfEventHandler(cacheHandler redis.CacheInterface, httpHandler httprest.RestInterface) ConfEventInterface {
	return &ConfEventHandler{
		cacheHandler: cacheHandler,
		httpHandler:  httpHandler,
	}
}

func (cf *ConfEventHandler) ProcessConfEvent(eventData []byte) error {
	var confEvent model.ConferenceEvent
	var err error
	var confCacheData []byte
	var statusCallbackUrl string
	var statusCallbackMethod = "POST"

	var confCacheModel model.ConferenceDetailsFromCache

	if err := json.Unmarshal(eventData, &confEvent); err == nil {

		log.Println("event is ", confEvent.EventName, " sub class is ", confEvent.EventSubclass, " action is ",
			confEvent.Action, " name is ", confEvent.ConferenceName, " time is ", confEvent.EventDateTimestamp)

		confKey := fmt.Sprintf("conference:%s@%s", url.PathEscape(confEvent.ConferenceName),
			confEvent.ConferenceProfileName)

		log.Println("conference get key is ", confKey)

		//getting details from conference cache
		if confCacheData, err = cf.cacheHandler.Get(confKey); err == nil {
			if err := json.Unmarshal(confCacheData, &confCacheModel); err == nil {
				statusCallbackUrl = confCacheModel.DialConferenceStatusCallback
				statusCallbackMethod = confCacheModel.DialConferenceStatusCallbackMethod
				log.Println("conference status callback url is ", confCacheModel)
			} else {
				log.Println("conference data found in cache, unmarshal failed ", confKey)
			}
		}

		if statusCallbackUrl != "" {
			statusCallbackMap := cf.FormatConferenceStatusCallback(confCacheModel, confEvent)

			if statusCallbackMethod == "GET" {
				_, _, _ = cf.httpHandler.Get(statusCallbackMap, statusCallbackUrl)
			} else {
				_, _, _ = cf.httpHandler.Post(statusCallbackMap, statusCallbackUrl)
			}
		}

	}
	return err
}

func (cf *ConfEventHandler) FormatConferenceStatusCallback(confCacheData model.ConferenceDetailsFromCache,
	confEventData model.ConferenceEvent) map[string]interface{} {
	var confEvent = make(map[string]interface{})

	confEvent["ConferenceSid"] = confCacheData.DialConfSid
	confEvent["FriendlyName"] = confEventData.ConferenceProfileName
	confEvent["AccountSid"] = confCacheData.DialConfAccountSid
	confEvent["StatusCallbackEvent"] = getConfEventStatus(confEventData.Action)
	confEvent["CallSid"] = confEventData.ChannelCallUUID
	return confEvent
}

func getConfEventStatus(confEventName string) string {
	switch confEventName {
	case "add-member":
		return "participant-join"
	case "del-member":
		return "participant-leave"
	case "conference_create":
		return "conference-start"
	case "conference-destroy":
		return "conference-end"
	default:
		return ""

	}
}
