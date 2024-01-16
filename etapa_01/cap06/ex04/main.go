package main

import "fmt"

func geraImpar() func() int {
	i := 1
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}

func main() {
	fmt.Println("Usando o `makeEvenGenerator` como exemplo, escreva uma função que genre números ímpares.")

	gerador := geraImpar()
	fmt.Println(gerador())
	fmt.Println(gerador())
	fmt.Println(gerador())
}
