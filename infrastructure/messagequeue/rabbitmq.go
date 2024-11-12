package messagequeue

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var (
	conn    *amqp.Connection
	channel *amqp.Channel
)

func ConnectRabbitMQ() error {
	var err error
	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return fmt.Errorf("failed to connect to rabbitmq: %v", err)
	}

	channel, err = conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}

	return nil
}

func CloseRabbitMQ() {
	if channel != nil {
		channel.Close()
	}
	if conn != nil {
		conn.Close()
	}
}

func DeclareExchange(name, kind, durable string) error {
	if err := channel.ExchangeDeclare(name, kind, durable, false, false, false, nil); err != nil {
		return fmt.Errorf("failed to declare exchange: %v", err)
	}
	return nil
}

func DeclareQueue(name, durable string) error {
	if err := channel.QueueDeclare(name, durable, false, false, false, false, nil); err != nil {
		return fmt.Errorf("failed to declare queue: %v", err)
	}
	return nil

}

func BindQueue(queue, exchange, routingKey string) error {
	if err := channel.QueueBind(queue, routingKey, exchange, false, nil); err != nil {
		return fmt.Errorf("failed to bind queue: %v", err)
	}
	return nil
}

func Publish(exchange, routingKey string, body []byte) error {
	if err := channel.Publish(exchange, routingKey, false, false, amqp.Publishing{
		Body: body,
	}); err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}
	return nil
}

func Consume(queue, exchange, routingKey string, autoAck bool) (<-chan amqp.Delivery, error) {
	messages := make(chan amqp.Delivery)

	if err := channel.QueueDeclare(queue, false, false, false, false, nil); err != nil {
		return nil, fmt.Errorf("failed to declare queue: %v", err)
	}

	if err := channel.QueueBind(queue, routingKey, exchange, false, nil); err != nil {
		return nil, fmt.Errorf("failed to bind queue: %v", err)
	}

	if err := channel.Consume(queue, "", autoAck, false, false, false, nil); err != nil {
		return nil, fmt.Errorf("failed to consume messages: %v", err)
	}

	return messages, nil
}