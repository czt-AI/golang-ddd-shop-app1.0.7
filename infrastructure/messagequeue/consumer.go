package messagequeue

import (
	"context"
	"log"

	"github.com/streadway/amqp"
)

type Consumer struct {
	client *RabbitMQClient
}

func NewConsumer(client *RabbitMQClient) *Consumer {
	return &Consumer{
		client: client,
	}
}

func (consumer *Consumer) Consumequeue, exchange, routingKey string, autoAck bool) (<-chan amqp.Delivery, error) {
	messages := make(chan amqp.Delivery)

	if err := consumer.client.Consume(queue, exchange, routingKey, autoAck, false, false, nil); err != nil {
		log.Printf("Failed to consume messages: %v", err)
		return nil, err
	}

	return messages, nil
}

func (consumer *Consumer) HandleMessage(messages <-chan amqp.Delivery) {
	for msg := range messages {
		log.Printf("Received message: %s", msg.Body)
		// Process the message
		// ...

		if msg.Acknowledged != nil {
			msg.Acknowledged(false)
		}
	}
}