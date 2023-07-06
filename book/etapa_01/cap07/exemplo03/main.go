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

type Circulo struct {
	x, y, r float64
}

func (c *Circulo) area() float64 {
	return math.Pi * c.r * c.r
}

type Retangulo struct {
	x1, y1, x2, y2 float64
}

func (r *Retangulo) area() float64 {
	l := dist(r.x1, r.y1, r.x1, r.y2)
	w := dist(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main() {

	meuCirculo := Circulo{0, 0, 5}
	fmt.Println("Área Cículo:", meuCirculo.area())

	meuRetangulo := Retangulo{0, 0, 10, 10}
	fmt.Println("Área Retângulo:", meuRetangulo.area())

}
