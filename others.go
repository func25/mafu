package mafu

func Sign[T SignedNumber](v T) T {
	zero := T(0)

	switch true {
	case v > zero:
		return 1
	case v < zero:
		return -1
	}

	return 0
}
