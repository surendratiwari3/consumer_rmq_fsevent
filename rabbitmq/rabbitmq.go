package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync"
	"time"
)

type consumerFunc func([]byte) error

type RmqInterface interface {
	PublishCallStats(data []byte) error
	Consumer(consumerTag string, consumerHandler consumerFunc) error
}

type RmqAdapter struct {
	url        string
	client     *amqp.Connection
	channel    *amqp.Channel
	routingKey string
	done       chan error
	mLock      sync.Mutex
}

// Connect opens a connection to RabbitMQ, declares an exchange, opens a channel,
// declares and binds the queue and enables publish notifications
func NewRmqAdapter(rmqURL, queue string) (RmqInterface, error) {

	var conn *amqp.Connection
	var channel *amqp.Channel
	var err error

	mutexLock := sync.Mutex{}

	if conn, err = amqp.Dial(rmqURL); err != nil {
		return nil, err
	}

	if channel, err = conn.Channel(); err != nil {
		return nil, err
	}

	rmqConn := &RmqAdapter{
		url:        rmqURL,
		client:     conn,
		channel:    channel,
		routingKey: queue,
		done:       make(chan error),
		mLock:      mutexLock,
	}

	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				rmqConn.reconnect()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return rmqConn, nil
}

func (cq *RmqAdapter) reconnect() {
	cq.mLock.Lock()
	if cq.client.IsClosed() {
		var conn *amqp.Connection
		var channel *amqp.Channel
		var err error
		_ = cq.channel.Close()
		if conn, err = amqp.Dial(cq.url); err == nil {
			cq.client = conn
			if channel, err = conn.Channel(); err != nil {
				//retry
			}
			cq.channel = channel
		}
	}
	cq.mLock.Unlock()
}

func (cq *RmqAdapter) Consumer(consumerTag string, consumerHandler consumerFunc) error {
	log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", consumerTag)
	deliveries, err := cq.channel.Consume(
		cq.routingKey, // name
		consumerTag,   // consumerTag,
		false,         // noAck
		false,         // exclusive
		false,         // noLocal
		false,         // noWait
		nil,           // arguments
	)
	if err != nil {
		return fmt.Errorf("Queue Consume: %s", err)
	}
	go cq.handle(deliveries, cq.done, consumerHandler)

	return err
}

func (cq *RmqAdapter) PublishCallStats(data []byte) error {
	var err error
	if err = cq.channel.Publish(
		"",
		cq.routingKey,
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            data,
			DeliveryMode:    amqp.Transient,
			Priority:        0,
		},
	); err != nil {
		return fmt.Errorf("Exchange Publish: %s", err)
	}
	return err
}

func (cq *RmqAdapter) handle(deliveries <-chan amqp.Delivery, done chan error, consumerHandler consumerFunc) {
	for d := range deliveries {
		if err := consumerHandler(d.Body); err == nil {
			d.Ack(false)
		}
	}
	log.Printf("handle: deliveries channel closed")
	done <- nil
}
