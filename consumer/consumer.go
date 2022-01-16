package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/consumer_rmq_fsevent/model"
	"github.com/streadway/amqp"
	"log"
)

type Consumer interface {
	Shutdown() error
}

type RabbitMqConsumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	tag     string
	done    chan error
}

func NewConsumer(amqpURI, exchange, exchangeType, queueName, key, cTag string) (Consumer, error) {
	c := &RabbitMqConsumer{
		conn:    nil,
		channel: nil,
		tag:     cTag,
		done:    make(chan error),
	}

	var err error

	log.Printf("dialing %q", amqpURI)
	c.conn, err = amqp.Dial(amqpURI)
	if err != nil {
		return nil, fmt.Errorf("Dial: %s", err)
	}

	go func() {
		fmt.Printf("closing: %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
	}()

	log.Printf("got Connection, getting Channel")
	c.channel, err = c.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Channel: %s", err)
	}

	log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", c.tag)
	deliveries, err := c.channel.Consume(
		queueName, // name
		c.tag,     // consumerTag,
		false,     // noAck
		false,     // exclusive
		false,     // noLocal
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Consume: %s", err)
	}

	go handle(deliveries, c.done)

	return c, nil
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

func handle(deliveries <-chan amqp.Delivery, done chan error) {
	for d := range deliveries {
		var confEvent model.ConferenceEvent
		if err := json.Unmarshal(d.Body, &confEvent); err == nil {
			log.Println("event is ", confEvent.EventName, " sub class is ",
				confEvent.EventSubclass, " action is ", confEvent.Action,
				" name is ", confEvent.ConferenceName, " time is ", confEvent.EventDateTimestamp)
			d.Ack(false)
		}
		/*log.Printf(
			"got %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)*/

	}
	log.Printf("handle: deliveries channel closed")
	done <- nil
}
