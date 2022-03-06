package pokemon

type TimesMarkedAsFavorite struct {
	value uint
}

func (timesMarkedAsFavorite TimesMarkedAsFavorite) GetValue() uint {
	return timesMarkedAsFavorite.value
}

func CreateTimesMarkedAsFavorite(value uint) TimesMarkedAsFavorite {
	return TimesMarkedAsFavorite{value}
}
