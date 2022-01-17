package event

import (
	"encoding/json"
	"fmt"
	"github.com/consumer_rmq_fsevent/httprest"
	"github.com/consumer_rmq_fsevent/model"
	"github.com/consumer_rmq_fsevent/redis"
	"log"
	"net/url"
	"strings"
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
		if confCacheData, err = cf.cacheHandler.Get(confKey); err != nil {
			log.Println("conference data not found in cache", confKey, err)
			return err
		}

		if err := json.Unmarshal(confCacheData, &confCacheModel); err != nil {
			log.Println("conference data found in cache, unmarshal failed ", confKey,err)
			return err
		}

		statusCallbackUrl = confCacheModel.DialConferenceStatusCallback

		statusCallbackMethod = confCacheModel.DialConferenceStatusCallbackMethod

		log.Println("conference status callback url is ", confCacheModel)

		if confEvent.Action == "conference-destroy" {
			if err = cf.cacheHandler.Expire(confKey); err != nil {
				log.Println("conference key delete failed on conference end ", confKey)
			}
		}

		if statusCallbackUrl != "" {
			log.Println("statuscallback url is ", statusCallbackUrl)
			statusCallbackMap := cf.FormatConferenceStatusCallback(confCacheModel, confEvent)
			log.Println("statuscallback map is before event check ", statusCallbackMap)
			if statusCallbackMap["StatusCallbackEvent"] != "" {
				log.Println("statuscallback map is ", statusCallbackMap)
				if statusCallbackMethod == "GET" {
					_, _, _ = cf.httpHandler.Get(statusCallbackMap, statusCallbackUrl)
				} else {
					_, _, _ = cf.httpHandler.Post(statusCallbackMap, statusCallbackUrl)
				}
			}
		}

	}
	return err
}

func (cf *ConfEventHandler) FormatConferenceStatusCallback(confCacheData model.ConferenceDetailsFromCache,
	confEventData model.ConferenceEvent) map[string]interface{} {
	var confEvent = make(map[string]interface{})
	confEvent["ConferenceSid"] = confCacheData.DialConfSid
	confEvent["FriendlyName"] = getConfFriendlyName(confEventData.ConferenceName)
	confEvent["AccountSid"] = confCacheData.DialConfAccountSid
	confEvent["StatusCallbackEvent"] = getConfEventStatus(confEventData.Action)
	confEvent["CallSid"] = confEventData.ChannelCallUUID
	return confEvent
}

func getConfFriendlyName(absoluteConfName string) string {
	return strings.SplitN(absoluteConfName, "-", 2)[1]
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
