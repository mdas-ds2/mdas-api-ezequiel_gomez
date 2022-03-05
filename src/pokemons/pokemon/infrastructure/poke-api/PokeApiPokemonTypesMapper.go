package pokeapi

import (
	"encoding/json"

	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/domain"
	sharedInfra "github.com/mdas-ds2/mdas-api-g3/src/shared/infrastructure"
)

type PokeApiResponse = []byte

func mapResponseToPokemon(response PokeApiResponse, timesMarkedAsFAvorite uint) (domain.Pokemon, error) {
	var pokemonResponse sharedInfra.PokemonModel
	json.Unmarshal(response, &pokemonResponse)

	pokeId := domain.CreatePokemonId(pokemonResponse.ID)
	name := domain.CreateName(pokemonResponse.Name)
	height := domain.CreateHeight(pokemonResponse.Height)
	weight := domain.CreateWeight(pokemonResponse.Weight)
	timesMarkedAsFavorite := domain.CreateTimesMarkedAsFavorite(timesMarkedAsFAvorite)

	pokemon := domain.CreatePokemon(pokeId, name, height, weight, timesMarkedAsFavorite)

	return pokemon, nil
}
