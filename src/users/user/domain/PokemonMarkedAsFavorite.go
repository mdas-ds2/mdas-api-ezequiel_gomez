package user

import (
	"time"
)

type PokemonMarkedAsFavorite struct {
	aggregateId string
	eventType   string
	createdAt   time.Time
}

func CreatePokemonMarkedAsFavorite(pokemonId PokemonId) PokemonMarkedAsFavorite {
	eventType := "pokemon-marked-as-favorite"
	return PokemonMarkedAsFavorite{pokemonId.GetValue(), eventType, time.Now()}
}

func (event PokemonMarkedAsFavorite) AggregateId() string {
	return event.aggregateId
}

func (event PokemonMarkedAsFavorite) EventType() string {
	return event.eventType
}

func (event PokemonMarkedAsFavorite) CreatedAt() time.Time {
	return event.createdAt
}
