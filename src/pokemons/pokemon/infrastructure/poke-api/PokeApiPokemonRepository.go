package pokeapi

import (
	"net/http"
	"strconv"

	shared "github.com/mdas-ds2/mdas-api-g3/src/shared/infrastructure"

	httpClient "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/http-client"
	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/domain"
)

var timesMarkedAsFavoriteHashTable = make(map[int]uint)

type PokeApiPokemonRepository struct{}

const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon/"

func (repository PokeApiPokemonRepository) Find(pokemonId domain.PokemonId) (domain.Pokemon, error) {
	pokemonIdStr := strconv.Itoa(pokemonId.GetValue())
	urlPath := pokeApiUrl + pokemonIdStr

	response, errorOnResponse := httpClient.Get(urlPath)
	timesMarkedAsFAvorite := timesMarkedAsFavoriteHashTable[pokemonId.GetValue()]

	if response.StatusCode == http.StatusServiceUnavailable {
		serviceUnavailableException := shared.CreatePokemonRepositoryUnavailableException()
		return domain.Pokemon{}, serviceUnavailableException.GetError()
	}

	if errorOnResponse != nil {
		return domain.Pokemon{}, errorOnResponse
	}

	if response.StatusCode == http.StatusNotFound {
		pokemonNotFoundException := domain.CreatePokemonNotFoundException(pokemonId)
		return domain.Pokemon{}, pokemonNotFoundException.GetError()
	}

	return mapResponseToPokemon(response.Body, timesMarkedAsFAvorite)
}

func (repository PokeApiPokemonRepository) Save(pokemon domain.Pokemon) error {
	pokemonId := pokemon.GetId().GetValue()
	timesMarkedAsFavorite := pokemon.GetTimesMarkedAsFavorite().GetValue()

	timesMarkedAsFavoriteHashTable[pokemonId] = timesMarkedAsFavorite

	return nil
}
