package pokemon

type Pokemon struct {
	id                    PokemonId
	name                  Name
	height                Height
	weight                Weight
	timesMarkedAsFavorite TimesMarkedAsFavorite
}

func CreatePokemon(id PokemonId, name Name, height Height, weight Weight, timesMarkedAsFavorite TimesMarkedAsFavorite) Pokemon {
	return Pokemon{id, name, height, weight, timesMarkedAsFavorite}
}

func (pokemon *Pokemon) GetId() PokemonId {
	return pokemon.id
}

func (pokemon *Pokemon) GetName() Name {
	return pokemon.name
}

func (pokemon *Pokemon) GetHeight() Height {
	return pokemon.height
}

func (pokemon *Pokemon) GetWeight() Weight {
	return pokemon.weight
}

func (pokemon *Pokemon) GetTimesMarkedAsFavorite() TimesMarkedAsFavorite {
	return pokemon.timesMarkedAsFavorite
}
