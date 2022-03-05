package user

import (
	shared "github.com/mdas-ds2/mdas-api-g3/src/shared/domain"
)

type User struct {
	shared.AggregateRoot
	id               UserId
	favoritePokemons PokemonIdCollection
}

func (user User) GetId() UserId {
	return user.id
}

func CreateUser(id UserId, favoritePokemons PokemonIdCollection) User {
	emptyEventCollection := shared.CreateDomainEventCollection([]shared.DomainEvent{})
	aggregateRoot := shared.CreateAggregateRoot(emptyEventCollection)

	return User{aggregateRoot, id, favoritePokemons}
}

func (user *User) AddFavorite(pokemonId PokemonId) error {
	if user.favoritePokemons.Has(pokemonId) {
		exception := CreateFavoritePokemonDuplicatedException(pokemonId)
		return exception.GetError()
	}

	user.favoritePokemons.Add(pokemonId)

	return nil
}

func (user *User) GetFavorites() PokemonIdCollection {
	return user.favoritePokemons
}
