package pokeapi

import (
	"encoding/json"

	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/domain"
	pokeApiShared "github.com/mdas-ds2/mdas-api-g3/src/shared/infrastructure"
)

type PokeApiResponse = []byte

func mapResponseToPokemon(response PokeApiResponse) (domain.Pokemon, error) {
	var pokemonResponse pokeApiShared.PokemonModel
	json.Unmarshal(response, &pokemonResponse)

	pokeId := domain.CreatePokemonId(pokemonResponse.ID)
	name := domain.CreateName(pokemonResponse.Name)
	height := domain.CreateHeight(pokemonResponse.Height)
	weight := domain.CreateWeight(pokemonResponse.Weight)
	timesMarkedAsFavorite := domain.CreateTimesMarkedAsFavorite(0)

	pokemon := domain.CreatePokemon(pokeId, name, height, weight, timesMarkedAsFavorite)

	return pokemon, nil
}
