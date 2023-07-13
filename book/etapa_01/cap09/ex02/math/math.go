package math

func Average(x []float64) float64 {
	if len(x) == 0 {
		return 0
	}

	total := 0.
	for _, v := range x {
		total += v
	}
	return total / float64(len(x))
}

func Min(x []float64) float64 {

	if len(x) == 0 {
		return 0
	}

	min := x[0]
	for _, v := range x {
		if min > v {
			min = v
		}
	}
	return min
}

func Max(x []float64) float64 {

	if len(x) == 0 {
		return 0
	}

	max := x[0]
	for _, v := range x {
		if max < v {
			max = v
		}
	}
	return max
}
