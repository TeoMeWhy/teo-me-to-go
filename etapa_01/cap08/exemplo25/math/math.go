package math

// Average é uma função para calcular média de uma fatia de float64
func Average(xs []float64) float64 {
	total := 0.
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}
