package shared

import "time"

type DomainEvent struct {
	aggregateId int
	eventType   string
	createdAt   time.Time
}

func CreateDomainEvent(aggregateId int, eventType string) DomainEvent {
	return DomainEvent{aggregateId, eventType, time.Now()}
}
