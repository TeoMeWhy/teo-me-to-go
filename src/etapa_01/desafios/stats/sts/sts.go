package sts

import "sort"

func SomaFloat64(x []float64) float64 {
	total := 0.
	for _, v := range x {
		total += v
	}
	return total
}

func AvgFloat64(x []float64) float64 {
	total := SomaFloat64(x)
	return float64(total) / float64(len(x))
}

func MinFloat64(x []float64) float64 {
	min := x[0]
	for _, v := range x {
		if min > v {
			min = v
		}
	}
	return min
}

func MaxFloat64(x []float64) float64 {
	max := x[0]
	for _, v := range x {
		if max < v {
			max = v
		}
	}
	return max
}

func MedianFloat64(x []float64) float64 {
	sort.Float64s(x)

	if len(x)%2 == 0 {
		return (x[len(x)/2] + x[len(x)/2+1]) / 2.
	} else {
		return x[len(x)/2+1]
	}

}
