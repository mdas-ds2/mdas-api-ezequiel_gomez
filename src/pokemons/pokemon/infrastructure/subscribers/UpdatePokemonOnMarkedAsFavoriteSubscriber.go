package subscribers

import (
	"encoding/json"
	"strconv"

	rabbitMQ "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/msg-broker/rabbitmq"
	application "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/application"
	pokeApi "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/infrastructure/poke-api"
	sharedInfra "github.com/mdas-ds2/mdas-api-g3/src/shared/infrastructure"
)

type UpdatePokemonOnMarkedAsFavoriteSubscriber struct {
	msgBus  rabbitMQ.AmqpBus
	useCase application.IncreaseTimesMarkedAsFavorite
}

func (subscriber UpdatePokemonOnMarkedAsFavoriteSubscriber) Run() {
	go func() {
		for d := range subscriber.msgBus {
			var payload sharedInfra.EventPayloadModel
			json.Unmarshal(d.Body, &payload)
			pokemonId, _ := strconv.Atoi(payload.AggregateId)

			subscriber.useCase.Execute(pokemonId, 1)
		}
	}()
}

func CreateUpdatePokemonOnMarkedAsFavoriteSubscriber(msgBus rabbitMQ.AmqpBus) UpdatePokemonOnMarkedAsFavoriteSubscriber {
	pokeApiPokemonRepository := pokeApi.PokeApiPokemonRepository{}
	useCase := application.IncreaseTimesMarkedAsFavorite{Repository: pokeApiPokemonRepository}

	return UpdatePokemonOnMarkedAsFavoriteSubscriber{msgBus, useCase}
}
