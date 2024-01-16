package main

import "fmt"

func main() {
	fmt.Println("Escreva um programa que descubra o menor número dessa lista:")

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	numero := x[0]

	for _, i := range x[1:] {
		if i < numero {
			numero = i
		}
	}

	fmt.Println("O menor número é:", numero)
}
