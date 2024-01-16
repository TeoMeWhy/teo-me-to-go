package main

import (
	"fmt"
	"math"
)

func dist(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func areaRetangulo(x1, y1, x2, y2 float64) float64 {
	l := dist(x1, y1, x1, y2)
	w := dist(x1, y1, x2, y1)
	return l * w
}

func areaCirculo(x, y, r float64) float64 {
	return math.Pi * r * r
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println("Area Retângulo:", areaRetangulo(rx1, ry1, rx2, ry2))
	fmt.Println("Área Círculo:", areaCirculo(cx, cy, cr))

}
