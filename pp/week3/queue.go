package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

type queue struct {
	channel *amqp.Channel
	conn    *amqp.Connection
}

func newQueue() *queue {
	connectRabbitMQ, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}

	_, err = channelRabbitMQ.QueueDeclare(
		"QueueService1", // queue name
		true,            // durable
		false,           // auto delete
		false,           // exclusive
		false,           // no wait
		nil,             // arguments
	)
	if err != nil {
		panic(err)
	}

	return &queue{
		channel: channelRabbitMQ,
		conn:    connectRabbitMQ,
	}
}

func (q *queue) publishEvent(event []byte) error {
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        event,
	}

	if err := q.channel.Publish(
		"",              // exchange
		"QueueService1", // queue name
		false,           // mandatory
		false,           // immediate
		message,         // message to publish
	); err != nil {
		return fmt.Errorf("publishing event to RabbitMQ: %v", err)
	}

	return nil
}

func (q *queue) consumeMessages() (<-chan amqp.Delivery, error) {
	messages, err := q.channel.Consume(
		"QueueService1", // queue name
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no local
		false,           // no wait
		nil,             // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("consuming messages: %v", err)
	}

	return messages, nil
}
