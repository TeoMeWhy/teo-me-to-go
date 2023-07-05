package main

import "fmt"

func main() {
	fmt.Print("Entre com o valor em pés para converter em metros: ")

	var p float64
	fmt.Scanf("%f", &p)

	m := p * 0.3048
	fmt.Println(m, "metros")
}
