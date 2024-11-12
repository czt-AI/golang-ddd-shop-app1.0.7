package messagequeue

import (
	"context"
	"log"

	"github.com/streadway/amqp"
)

type Producer struct {
	client *RabbitMQClient
}

func NewProducer(client *RabbitMQClient) *Producer {
	return &Producer{
		client: client,
	}
}

func (producer *Producer) Publishexchange, routingKey string, body []byte) error {
	if err := producer.client.Publish(exchange, routingKey, body); err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}
	return nil
}