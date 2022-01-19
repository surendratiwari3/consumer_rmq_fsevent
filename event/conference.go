package event

import (
	"encoding/json"
	"fmt"
	"github.com/consumer_rmq_fsevent/httprest"
	"github.com/consumer_rmq_fsevent/model"
	"github.com/consumer_rmq_fsevent/redis"
	"log"
	"strconv"
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

type FsEventAction string

// ConferenceStatus Enumerations
const (
	StatusConferenceEnd          ConferenceStatus = "conference-end"
	StatusConferenceStart        ConferenceStatus = "conference-start"
	StatusConferenceCreate       ConferenceStatus = "conference-create"
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

const (
	FsActionConferenceCreate FsEventAction = "conference_create"
	FsActionConferenceEnd    FsEventAction = "conference-destroy"
	FsActionAddMember        FsEventAction = "add-member"
	FsActionDelMember        FsEventAction = "del-member"
)

var ConfEventActionStatusMap = map[FsEventAction]ConferenceStatus{
	FsActionConferenceCreate: StatusConferenceCreate,
	FsActionConferenceEnd:    StatusConferenceEnd,
	FsActionAddMember:        StatusParticipantJoin,
	FsActionDelMember:        StatusParticipantLeave,
}

func NewConfEventHandler(cacheHandler redis.CacheInterface, httpHandler httprest.RestInterface) ConfEventInterface {
	return &ConfEventHandler{
		cacheHandler: cacheHandler,
		httpHandler:  httpHandler,
	}
}

//convert json byte to fsevent
func (cf *ConfEventHandler) ProcessConfEvent(eventData []byte) error {
	var confFsEvent model.ConferenceFsEvent
	var err error

	if err := json.Unmarshal(eventData, &confFsEvent); err == nil {

		log.Println("event is ", confFsEvent.EventName, " sub class is ", confFsEvent.EventSubclass, " action is ",
			confFsEvent.Action, " name is ", confFsEvent.ConferenceName, " time is ", confFsEvent.EventDateTimestamp)

		//no need to send this to status http url
		if confFsEvent.StatuscallbackURL == "" {
			return nil
		}
		log.Println("event_data - ", confFsEvent)

		return cf.ProcessFsEventToStatusCallback(confFsEvent)

	}
	return err
}

//processing fsevent to statuscallback and send to url
func (cf *ConfEventHandler) ProcessFsEventToStatusCallback(fsEvent model.ConferenceFsEvent) error {
	var sequenceNumber int64
	var err error
	var confCommonModel model.ConferenceCommon

	confKey := fmt.Sprintf("conference:sequence:%s", fsEvent.ConferenceUniqueID)

	if sequenceNumber, err = cf.cacheHandler.Incr(confKey); err != nil {
		log.Println("error while getting conference sequence number for ", confKey, err)
	}

	confCommonModel.SequenceNumber = strconv.Itoa(int(sequenceNumber))

	confCommonModel.StatusCallbackEvent = string(ConfEventActionStatusMap[FsEventAction(fsEvent.Action)])

	confCommonModel.Timestamp = fsEvent.EventDateGmt

	confCommonModel.ConferenceSid = fsEvent.ConferenceUniqueID

	confCommonModel.AccountSid, confCommonModel.FriendlyName = cf.getConfFriendlyName(fsEvent.ConferenceName)

	//expire the sequence number key, later we need to do it based on conference side
	if fsEvent.Action == "conference-destroy" {
		if err = cf.cacheHandler.Expire(confKey); err != nil {
			log.Println("conference key delete failed on conference end ", confKey)
		}
	}

	//if event name is nil, do not send the event
	if confCommonModel.StatusCallbackEvent == "" {
		log.Println("conference event name is missing, not sending statuscallback ")
		return nil
	}

	dataMap := make(map[string]interface{})
	if callbackByte, err := json.Marshal(confCommonModel); err == nil {
		if err := json.Unmarshal(callbackByte, &dataMap); err != nil {
			log.Println("unmarshal failed on convert statuscallback to map", err)
			return err
		}
	}

	if fsEvent.StatuscallbackMethod == "GET" || fsEvent.StatuscallbackMethod == "get" {
		_, _, err = cf.httpHandler.Get(dataMap, fsEvent.StatuscallbackURL)
		return err
	}
	_, _, err = cf.httpHandler.Post(dataMap, fsEvent.StatuscallbackURL)

	return err
}

func (cf *ConfEventHandler) getConfFriendlyName(absoluteConfName string) (string, string) {
	authConf := strings.SplitN(absoluteConfName, "-", 2)
	return authConf[0], authConf[1]
}
