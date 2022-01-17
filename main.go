// This declares a durable Exchange, an ephemeral (auto-delete) Queue,
// binds the Queue to the Exchange with a binding key, and consumes every
// message published to that Exchange with that routing key.
//
package main

import (
	"flag"
	"github.com/consumer_rmq_fsevent/event"
	"github.com/consumer_rmq_fsevent/httprest"
	"github.com/consumer_rmq_fsevent/rabbitmq"
	"github.com/consumer_rmq_fsevent/redis"
	"log"
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

	/*rmqReq := consumer.RmqConsumerRequest{
		AmqpURI:      *uri,
		QueueName:    *queue,
		Ctag:         *consumerTag,
		ConfEventHandler:confEventHandler,
	}*/

	if rmqHandle, err := rabbitmq.NewRmqAdapter(*uri, *queue); err != nil {
		log.Fatalf("%s", err)
	} else {
		if err := rmqHandle.Consumer(*consumerTag, confEventHandler.ProcessConfEvent); err != nil {
			log.Fatalf("%s", err)
		} else {
			select {}
		}
	}
}
