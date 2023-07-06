package main

import "fmt"

func swap(x, y *int) {

	*x, *y = *y, *x

}

func main() {
	fmt.Println("Escreva um programa que possa trocar dois inteiros (x:=1; y:=2; swap(&x, &y)) deve resultar em x=2 e y=1.")

	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)
}
