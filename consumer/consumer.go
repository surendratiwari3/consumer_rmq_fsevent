package consumer

import (
	"fmt"
	"github.com/consumer_rmq_fsevent/event"
	"github.com/streadway/amqp"
	"log"
)

type Consumer interface {
	Shutdown() error
}

type RmqConsumerRequest struct {
	AmqpURI string
	QueueName string
	Ctag string
	ConfEventHandler event.ConfEventInterface
}

type RabbitMqConsumer struct {
	conn    *amqp.Connection
	rmqReq RmqConsumerRequest
	channel *amqp.Channel
	tag     string
	done    chan error
}

func NewConsumer(rmqRequest RmqConsumerRequest) (error) {
	c := &RabbitMqConsumer{
		conn:    nil,
		channel: nil,
		rmqReq:rmqRequest,
		tag:     rmqRequest.Ctag,
		done:    make(chan error),
	}

	var err error

	log.Printf("dialing %q", rmqRequest.AmqpURI)
	c.conn, err = amqp.Dial(rmqRequest.AmqpURI)
	if err != nil {
		return fmt.Errorf("Dial: %s", err)
	}

	go func() {
		fmt.Printf("closing: %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
	}()

	log.Printf("got Connection, getting Channel")
	c.channel, err = c.conn.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}

	log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", c.tag)
	deliveries, err := c.channel.Consume(
		rmqRequest.QueueName, // name
		c.tag,     // consumerTag,
		false,     // noAck
		false,     // exclusive
		false,     // noLocal
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("Queue Consume: %s", err)
	}

	go c.handle(deliveries, c.done)

	return nil
}

func (c *RabbitMqConsumer) Shutdown() error {
	// will close() the deliveries channel
	if err := c.channel.Cancel(c.tag, true); err != nil {
		return fmt.Errorf("Consumer cancel failed: %s", err)
	}

	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	defer log.Printf("AMQP shutdown OK")

	// wait for handle() to exit
	return <-c.done
}

func (ch *RabbitMqConsumer) handle(deliveries <-chan amqp.Delivery, done chan error) {
	for d := range deliveries {
		if err := ch.rmqReq.ConfEventHandler.ProcessConfEvent(d.Body); err == nil{
			d.Ack(false)
		}
	}
	log.Printf("handle: deliveries channel closed")
	done <- nil
}
