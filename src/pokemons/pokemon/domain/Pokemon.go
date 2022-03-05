package pokemon

import (
	shared "github.com/mdas-ds2/mdas-api-g3/src/shared/domain"
)

type Pokemon struct {
	shared.AggregateRoot
	id                    PokemonId
	name                  Name
	height                Height
	weight                Weight
	timesMarkedAsFavorite TimesMarkedAsFavorite
}

func CreatePokemon(pokemonId PokemonId, name Name, height Height, weight Weight, timesMarkedAsFavorite TimesMarkedAsFavorite) Pokemon {
	emptyEventCollection := shared.CreateDomainEventCollection([]shared.DomainEvent{})
	aggregateRoot := shared.CreateAggregateRoot(emptyEventCollection)

	return Pokemon{aggregateRoot, pokemonId, name, height, weight, timesMarkedAsFavorite}
}

func (pokemon *Pokemon) GetId() PokemonId {
	return pokemon.id
}

func (pokemon *Pokemon) GetName() Name {
	return pokemon.name
}

func (pokemon *Pokemon) GetHeight() Height {
	return pokemon.height
}

func (pokemon *Pokemon) GetWeight() Weight {
	return pokemon.weight
}

func (pokemon *Pokemon) GetTimesMarkedAsFavorite() TimesMarkedAsFavorite {
	return pokemon.timesMarkedAsFavorite
}
