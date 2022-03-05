package pokemon

type Pokemon struct {
	id                    Id
	name                  Name
	height                Height
	weight                Weight
	timesMarkedAsFavorite TimesMarkedAsFavorite
}

func CreatePokemon(id Id, name Name, height Height, weight Weight, timesMarkedAsFavorite TimesMarkedAsFavorite) Pokemon {
	return Pokemon{id, name, height, weight, timesMarkedAsFavorite}
}

func (pokemon *Pokemon) GetId() Id {
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
