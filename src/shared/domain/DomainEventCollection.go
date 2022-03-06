package shared

type DomainEventCollection struct {
	events []DomainEvent
}

func (collection *DomainEventCollection) Add(event DomainEvent) {
	collection.events = append(collection.events, event)
}

func (collection *DomainEventCollection) PullAll() []DomainEvent {
	var events []DomainEvent
	events = append(events, collection.events...)

	collection.events = []DomainEvent{}

	return events
}

func CreateDomainEventCollection(events []DomainEvent) DomainEventCollection {
	eventsSlice := []DomainEvent{}
	return DomainEventCollection{eventsSlice}
}
