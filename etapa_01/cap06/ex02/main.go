package main

import "fmt"

func par(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func main() {
	fmt.Println("Escreva uma função que aceite um inteiro, calcule sua metade e devolva verdadeiro, caso o número seja par e falso se for ímpar.")

	numero := 0
	fmt.Scanf("%d", &numero)
	fmt.Println(par(numero))
}
