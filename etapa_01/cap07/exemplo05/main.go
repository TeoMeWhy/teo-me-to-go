package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	x, y, r float64
}

func (c *Circulo) area() float64 {
	return math.Pi * c.r * c.r
}

func AreaTotal(c ...Circulo) float64 {
	total := 0.
	for _, v := range c {
		total += v.area()
	}
	return total
}

func main() {
	c1 := Circulo{0, 0, 10}
	c2 := Circulo{0, 0, 10}
	fmt.Println(AreaTotal(c1, c2))
}
