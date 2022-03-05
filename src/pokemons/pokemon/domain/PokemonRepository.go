package pokemon

type PokemonRepository interface {
	Find(id PokemonId) (Pokemon, error)
	Save(Pokemon) error
}
