package messagequeue

import (
	"context"
	"log"

	"github.com/streadway/amqp"
)

type Handler func(context.Context, amqp.Delivery) error

var handlers = make(map[string]Handler)

func RegisterHandler(queue string, handler Handler) {
	handlers[queue] = handler
}

func Consume(queue string, exchange, routingKey string, autoAck bool) error {
	messages, err := rmqClient.Consume(queue, exchange, routingKey, autoAck, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for msg := range messages {
			log.Printf("Received message: %s", msg.Body)

			if handler, ok := handlers[queue]; ok {
				if err := handler(context.Background(), msg); err != nil {
					log.Printf("Handler failed: %v", err)
					if autoAck {
						msg.Acknowledged(false)
					}
				} else if autoAck {
					msg.Acknowledged(true)
				}
			} else {
				log.Printf("No handler registered for queue: %s", queue)
				if autoAck {
					msg.Acknowledged(false)
				}
			}
		}
	}()

	return nil
}