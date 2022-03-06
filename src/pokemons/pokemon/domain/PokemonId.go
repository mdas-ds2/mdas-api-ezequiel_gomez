package pokemon

type PokemonId struct {
	value int
}

func (id PokemonId) GetValue() int {
	return id.value
}

func CreatePokemonId(value int) PokemonId {
	return PokemonId{value}
}
