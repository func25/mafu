package mafu

func Sign[T Number](v T) int {
	zero := T(0)

	switch true {
	case v > zero:
		return 1
	case v < zero:
		return -1
	}

	return 0
}
