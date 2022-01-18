package main

import (
	"flag"
	"github.com/consumer_rmq_fsevent/event"
	"github.com/consumer_rmq_fsevent/httprest"
	"github.com/consumer_rmq_fsevent/rabbitmq"
	"github.com/consumer_rmq_fsevent/redis"
	"log"
	"os"
)

var (
	uri         = flag.String("uri", "amqp://user4tiniyo:4pass4tiniyo@3.0.39.201:5672/", "AMQP URI")
	queue       = flag.String("queue", "call_queue_stats", "Ephemeral AMQP queue name")
	consumerTag = flag.String("consumer-tag", "simple-consumer", "AMQP consumer tag (should not be blank)")
)

func init() {
	flag.Parse()
}

func main() {
	httpHandler := httprest.NewHttpRestHandler()

	cacheHandle, _ := redis.NewCacheHandler()

	confEventHandler := event.NewConfEventHandler(cacheHandle, httpHandler)

	var rabbitMqHandle rabbitmq.RmqInterface
	var err error

	if rabbitMqHandle, err = rabbitmq.NewRmqAdapter(*uri, *queue); err != nil {
		log.Fatalf("%s", err)
		os.Exit(0)
	}

	errConsumerChan := make(chan error)
	if err := rabbitMqHandle.Consumer(*consumerTag, confEventHandler.ProcessConfEvent, errConsumerChan); err != nil {
		log.Fatalf("%s", err)
		os.Exit(0)
	}

	err = <-errConsumerChan
	if err != nil {
		os.Exit(0)
	}
}
