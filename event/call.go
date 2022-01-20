package event

import (
	"encoding/json"
	"fmt"
	"github.com/consumer_rmq_fsevent/httprest"
	"github.com/consumer_rmq_fsevent/model"
	"github.com/consumer_rmq_fsevent/redis"
	jsoniter "github.com/json-iterator/go"
	"log"
	"strconv"
)

var fsEventCallStatus = map[string]string{
	"CHANNEL_ORIGINATE":"initiated",
	"CHANNEL_PROGRESS_MEDIA":"ringing",
	"CHANNEL_ANSWER":"answered",
	"CHANNEL_HANGUP_COMPLETE":"completed",
}

var freeswitchJson = jsoniter.Config{
	EscapeHTML:             true,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
	TagKey:                 "freeswitch",
}.Froze()

type CallEventInterface interface {
	ProcessCallEvent(eventData []byte) error
}

type CallEventHandler struct {
	httpHandler  httprest.RestInterface
	cacheHandler redis.CacheInterface
}

func NewCallEventHandler(cacheHandler redis.CacheInterface, httpHandler httprest.RestInterface) CallEventInterface {
	return &CallEventHandler{
		cacheHandler: cacheHandler,
		httpHandler:  httpHandler,
	}
}

//convert json byte to fsevent
func (cf *CallEventHandler) ProcessCallEvent(eventData []byte) error {
	var callFsEvent model.FsCallEvent
	var err error

	log.Println("event_data - ", string(eventData))

	if err := freeswitchJson.Unmarshal(eventData, &callFsEvent); err == nil {

		//no need to send this to status http url
		if callFsEvent.StatuscallbackURL == "" {
			return nil
		}
		log.Println("event_data - ", callFsEvent)

		return cf.ProcessFsEventToStatusCallback(callFsEvent)

	}
	return err
}

//processing fsevent to statuscallback and send to url
func (cf *CallEventHandler) ProcessFsEventToStatusCallback(fsEvent model.FsCallEvent) error {
	var sequenceNumber int64
	var err error

	confKey := fmt.Sprintf("call:sequence:%s", fsEvent.CallSid)

	fsEvent.CallStatus = fsEventCallStatus[fsEvent.EventName]

	//if event name is nil, do not send the event
	if fsEvent.CallStatus == "" {
		log.Println("conference event name is missing, not sending statuscallback ")
		return nil
	}

	if sequenceNumber, err = cf.cacheHandler.Incr(confKey); err != nil {
		log.Println("error while getting conference sequence number for ", confKey, err)
	}

	fsEvent.SequenceNumber = strconv.Itoa(int(sequenceNumber))

	fsEvent.ApiVersion = "v1"

	fsEvent.CallbackSource = "call-progress-events"

	dataMap := make(map[string]interface{})
	if callbackByte, err := json.Marshal(fsEvent); err == nil {
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