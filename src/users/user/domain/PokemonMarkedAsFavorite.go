package user

import (
	shared "github.com/mdas-ds2/mdas-api-g3/src/shared/domain"
)

type PokemonMarkedAsFavorite struct {
	shared.DomainEvent
}

func CreatePokemonMarkedAsFavorite(pokemonId PokemonId) PokemonMarkedAsFavorite {
	eventType := "pokemon-marked-as-favorite"
	domainEvent := shared.CreateDomainEvent(pokemonId.GetValue(), eventType)
	return PokemonMarkedAsFavorite{domainEvent}
}
