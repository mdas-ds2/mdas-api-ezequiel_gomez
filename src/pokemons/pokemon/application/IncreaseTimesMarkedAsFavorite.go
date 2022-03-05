package pokemon

import (
	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/domain"
)

type IncreaseTimesMarkedAsFavorite struct {
	Repository domain.PokemonRepository
}

func (useCase IncreaseTimesMarkedAsFavorite) Execute(pokemonId int, times uint) error {
	id := domain.CreatePokemonId(pokemonId)
	pokemon, err := useCase.Repository.Find(id)

	if err != nil {
		pokemonNotFoundException := domain.CreatePokemonNotFoundException(id)
		return pokemonNotFoundException.GetError()
	}

	pokemon.IncreaseTimesMarkedAsFavorite(times)

	useCase.Repository.Save(pokemon)

	return nil
}
