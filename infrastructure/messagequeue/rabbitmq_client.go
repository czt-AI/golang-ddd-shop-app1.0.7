package messagequeue

import (
	"context"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQClient() (*RabbitMQClient, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to rabbitmq: %v", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %v", err)
	}

	return &RabbitMQClient{
		conn:    conn,
		channel: channel,
	}, nil
}

func (client *RabbitMQClient) Close() {
	if client.channel != nil {
		client.channel.Close()
	}
	if client.conn != nil {
		client.conn.Close()
	}
}

func (client *RabbitMQClient) DeclareExchange(name, kind, durable string) error {
	if err := client.channel.ExchangeDeclare(name, kind, durable, false, false, false, nil); err != nil {
		return fmt.Errorf("failed to declare exchange: %v", err)
	}
	return nil
}

func (client *RabbitMQClient) DeclareQueue(name, durable string) error {
	if err := client.channel.QueueDeclare(name, durable, false, false, false, false, nil); err != nil {
		return fmt.Errorf("failed to declare queue: %v", err)
	}
	return nil

}

func (client *RabbitMQClient) BindQueue(queue, exchange, routingKey string) error {
	if err := client.channel.QueueBind(queue, routingKey, exchange, false, nil); err != nil {
		return fmt.Errorf("failed to bind queue: %v", err)
	}
	return nil
}

func (client *RabbitMQClient) Publish(exchange, routingKey string, body []byte) error {
	if err := client.channel.Publish(exchange, routingKey, false, false, amqp.Publishing{
		Body: body,
	}); err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}
	return nil
}

func (client *RabbitMQClient) Consume(queue, exchange, routingKey string, autoAck bool) (<-chan amqp.Delivery, error) {
	messages := make(chan amqp.Delivery)

	if err := client.channel.QueueDeclare(queue, false, false, false, false, nil); err != nil {
		return nil, fmt.Errorf("failed to declare queue: %v", err)
	}

	if err := client.channel.QueueBind(queue, routingKey, exchange, false, nil); err != nil {
		return nil, fmt.Errorf("failed to bind queue: %v", err)
	}

	if err := client.channel.Consume(queue, "", autoAck, false, false, false, nil); err != nil {
		return nil, fmt.Errorf("failed to consume messages: %v", err)
	}

	return messages, nil
}