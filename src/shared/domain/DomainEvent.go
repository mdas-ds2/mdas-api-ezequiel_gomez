package shared

import "time"

type DomainEvent interface {
	AggregateId() string
	EventType() string
	CreatedAt() time.Time
}
