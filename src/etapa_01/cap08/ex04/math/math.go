package math

func Average(x []float64) float64 {
	total := 0.
	for _, v := range x {
		total += v
	}

	return total / float64(len(x))
}

func Min(x []float64) float64 {
	min := x[0]
	for _, v := range x {
		if min > v {
			min = v
		}
	}
	return min
}

func Max(x []float64) float64 {
	max := x[0]
	for _, v := range x {
		if max < v {
			max = v
		}
	}
	return max
}
