package shared

type AggregateRoot struct {
	events DomainEventCollection
}

func (aggregate *AggregateRoot) Raise(event DomainEvent) {
	aggregate.events.Add(event)
}

func (aggregate *AggregateRoot) PullDomainEvents() []DomainEvent {
	return aggregate.events.PullAll()
}

func CreateAggregateRoot(events DomainEventCollection) AggregateRoot {
	return AggregateRoot{events}
}
