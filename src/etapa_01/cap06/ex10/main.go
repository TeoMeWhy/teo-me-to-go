package main

import "fmt"

func square(x *float64) {
	*x = *x * *x
}

func main() {
	fmt.Println("Qual e o valor de x após a execução deste programa?")

	x := 1.5

	square(&x)

	fmt.Println(x)
}
