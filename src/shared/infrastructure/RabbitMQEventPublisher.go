package shared

import (
	"fmt"

	domain "github.com/mdas-ds2/mdas-api-g3/src/shared/domain"
)

type RabbitMQEventPublisher struct{}

func (publisher RabbitMQEventPublisher) Publish(events []domain.DomainEvent) {
	for _, event := range events {
		publisher.sendEvent(event)
	}
}

func (publisher RabbitMQEventPublisher) sendEvent(event domain.DomainEvent) {
	// TODO: Convert to RMQ event and send it
	fmt.Println("Event sent", event.EventType())
}

func CreateRabbitMQEventPublisher() RabbitMQEventPublisher {
	return RabbitMQEventPublisher{}
}
