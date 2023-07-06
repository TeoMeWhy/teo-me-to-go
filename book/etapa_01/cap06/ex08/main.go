package main

import "fmt"

func main() {
	fmt.Println("Como podemos atribuir um valor a um ponteiro?")

	fmt.Println("*xPtr = 1")

	x := 0

	xPtr := &x

	*xPtr = 10

	fmt.Println(x)
}
