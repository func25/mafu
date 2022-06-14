package mafu

import "math"

func MinI(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func MinIE(first int, arr ...int) int {
	min := first
	for _, v := range arr {
		if v < min {
			min = v
		}
	}

	return min
}

func MinF(a float64, b float64) float64 {
	return math.Min(a, b)
}

func MaxI(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func MaxF(a float64, b float64) float64 {
	return math.Max(a, b)
}
