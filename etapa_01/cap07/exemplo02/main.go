package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	x, y, r float64
}

func areaCirculo(c *Circulo) float64 {
	return math.Pi * c.r * c.r
}

func main() {

	meuCirculo := Circulo{x: 0, y: 0, r: 5}
	fmt.Println(meuCirculo)
	fmt.Println(meuCirculo.x)
	fmt.Println(meuCirculo.y)
	fmt.Println(meuCirculo.r)

	meuCirculo.x = 10
	meuCirculo.y = 15

	fmt.Println(meuCirculo)

	fmt.Println("Área Círculo:", areaCirculo(&meuCirculo))
}
