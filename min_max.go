package mafu

func Min[T Ordered](first T, arr ...T) T {
	min := first
	for i := range arr {
		if arr[i] < min {
			min = arr[i]
		}
	}

	return min
}

func Max[T Ordered](first T, arr ...T) T {
	max := first
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}
