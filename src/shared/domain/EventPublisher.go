package shared

type EventPublisher interface {
	Publish(events DomainEvent)
}
