package subscribers

import (
	"strconv"

	application "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/application"
	sharedDomain "github.com/mdas-ds2/mdas-api-g3/src/shared/domain"
)

type UpdatePokemonOnMarkedAsFavoriteSubscriber struct {
	useCase application.IncreaseTimesMarkedAsFavorite
}

func (subscriber UpdatePokemonOnMarkedAsFavoriteSubscriber) On(event sharedDomain.DomainEvent) {
	pokemonId, _ := strconv.Atoi(event.AggregateId())

	subscriber.useCase.Execute(pokemonId, 1)
}

func CreateUpdatePokemonOnMarkedAsFavoriteSubscriber() UpdatePokemonOnMarkedAsFavoriteSubscriber {
	return UpdatePokemonOnMarkedAsFavoriteSubscriber{}
}
