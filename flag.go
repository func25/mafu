package mafu

//flag is used just for convenient, if you want to safely use this feature, please utilize "math/bits"

type Flag uint

func (b Flag) On(other Flag) Flag {
	return FlagOn(b, other)
}

func (b Flag) Off(other Flag) Flag {
	return FlagOff(b, other)
}

func (b Flag) Has(other Flag) bool {
	return FlagHas(b, other)
}

func FlagOn[T Integer](src T, other T) T {
	return src | other
}

func FlagOff[T Integer](src T, other T) T {
	return src &^ other
}

func FlagHas[T Integer](src T, other T) bool {
	return (src & other) == other
}
