package user

type User struct {
	id               UserId
	favoritePokemons PokemonIdCollection
}

func (user User) GetId() UserId {
	return user.id
}

func CreateUser(id UserId, favoritePokemons PokemonIdCollection) *User {
	user := &User{
		id:               id,
		favoritePokemons: favoritePokemons,
	}

	return user
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
