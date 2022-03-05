package pokemon

type Repository interface {
	Find(id PokemonId) (Pokemon, error)
}
