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
	return c.r * c.r * math.Pi
}

func (c *Circulo) perimetro() float64 {
	return 2 * math.Pi * c.r
}

type Retangulo struct {
	x1, y1, x2, y2 float64
}

func (r *Retangulo) area() float64 {
	l := dist(r.x1, r.y1, r.x1, r.y2)
	w := dist(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Retangulo) perimetro() float64 {
	l := dist(r.x1, r.y1, r.x1, r.y2)
	w := dist(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

type Shape interface {
	area() float64
	perimetro() float64
}

func main() {

	c1 := Circulo{0, 0, 10}
	c2 := Circulo{0, 0, 15}
	r1 := Retangulo{0, 0, 1, 10}

	shapes := []Shape{&c1, &c2, &r1}

	for _, s := range shapes {
		fmt.Println(s.area())
	}

}
