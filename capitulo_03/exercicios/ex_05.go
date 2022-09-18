package main

import "fmt"

func main() {

	fmt.Println("Ex 05: Usando o programa de exemplo como ponto de partida, escreva um programa que converta Fahrenheit em Celsius (C = (F - 32) * 5 / 9)")

	var fahrenheit float64

	fmt.Println("Entre com um valor em Fahrenheit a ser convertido para Celsius: ")
	fmt.Scanf("%f", &fahrenheit)

	celsius := (fahrenheit - 32.0) * (5.0 / 9.0)
	fmt.Println(fahrenheit, "Fahrenheit =", celsius, "celsius")

}
