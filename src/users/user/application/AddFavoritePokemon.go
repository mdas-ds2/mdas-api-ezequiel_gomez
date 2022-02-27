package user

import (
	domain "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

type AddFavoritePokemon struct {
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

	return nil
}
