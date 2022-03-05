package pokemon

type TimesMarkedAsFavorite struct {
	value int
}

func (timesMarkedAsFavorite TimesMarkedAsFavorite) GetValue() int {
	return timesMarkedAsFavorite.value
}

func New(value int) TimesMarkedAsFavorite {
	return TimesMarkedAsFavorite{value}
}
