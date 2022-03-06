package shared

type EventSubscriber interface {
	On(event DomainEvent)
}
