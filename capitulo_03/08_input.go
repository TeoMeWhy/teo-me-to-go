package main

import "fmt"

func main() {
	var number float64
	fmt.Println("Entre com um número: ")

	fmt.Scanf("%f", &number)

	output := number * 2
	fmt.Println(output)
}
