package mafu

import "math"

func MinI(a int, b int) int {
	if a < b {
		return a
	}

	return b
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
