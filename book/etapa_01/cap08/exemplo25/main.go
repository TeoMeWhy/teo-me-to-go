package main

import (
	"exemplo25/math"
	"fmt"
)

func main() {
	x := []float64{6.9, 9.10, 5.67, 10.}

	media := math.Average(x)
	fmt.Println(media)
}
