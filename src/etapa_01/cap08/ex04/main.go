package main

import (
	"ex04/math"
	"fmt"
)

func main() {
	x := []float64{4, 6.7, 5.8, 3.8, 9.2, 10}

	media := math.Average(x)
	fmt.Println("Média:", media)

	min := math.Min(x)
	fmt.Println("Min:", min)

	max := math.Max(x)
	fmt.Println("Max:", max)
}
