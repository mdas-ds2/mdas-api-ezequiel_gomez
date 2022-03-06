package user

import (
	sharedDomain "github.com/mdas-ds2/mdas-api-g3/src/shared/domain"
	domain "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

type AddFavoritePokemon struct {
	Publisher  sharedDomain.EventPublisher
	Repository domain.UserRepository
}

func (useCase AddFavoritePokemon) Execute(userId, pokemonId string) error {
	user := useCase.Repository.Find(
		domain.CreateUserId(userId),
	)

	err := user.AddFavorite(
		domain.CreatePokemonId(pokemonId),
	)

	if err != nil {
		return err
	}

	useCase.Repository.Save(user)

	domainEvents := user.PullDomainEvents()

	useCase.Publisher.Publish(domainEvents)

	return nil
}
