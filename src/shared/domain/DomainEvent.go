package shared

import "time"

type DomainEvent struct {
	aggregateId interface{}
	eventType   string
	createdAt   time.Time
}

func CreateDomainEvent(aggregateId interface{}, eventType string) DomainEvent {
	return DomainEvent{aggregateId, eventType, time.Now()}
}
