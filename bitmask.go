package mafu

//UIntBits is used just for convenient, if you want to safely use this feature, please utilize "math/bits"
type UIntBits uint

func (b UIntBits) On(other UIntBits) UIntBits {
	return b | other
}

func (b *UIntBits) SetOn(other UIntBits) {
	*b |= other
}

func (b UIntBits) Off(other UIntBits) UIntBits {
	return b &^ other
}

func (b *UIntBits) SetOff(other UIntBits) {
	*b = *b &^ other
}

func (b UIntBits) Has(other UIntBits) bool {
	return (b & other) == other
}
