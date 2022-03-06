package shared

import (
	"encoding/json"
	"fmt"

	rabbitMQ "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/msg-broker/rabbitmq"
	domain "github.com/mdas-ds2/mdas-api-g3/src/shared/domain"
)

type RabbitMQEventPublisher struct {
	msgQueue rabbitMQ.RabbitQueue
}

func (publisher RabbitMQEventPublisher) Publish(events []domain.DomainEvent) {
	for _, event := range events {
		publisher.sendEvent(event)
	}
}

func (publisher RabbitMQEventPublisher) sendEvent(event domain.DomainEvent) {
	fmt.Println("Event sent", event.EventType())

	bodyJson, _ := json.Marshal(map[string]string{
		"aggregate": event.AggregateId(),
		"eventType": event.EventType(),
	})

	publisher.msgQueue.Publish(bodyJson)
}

func CreateRabbitMQEventPublisher(msgQueue rabbitMQ.RabbitQueue) RabbitMQEventPublisher {
	return RabbitMQEventPublisher{msgQueue}
}
