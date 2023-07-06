package main

import "fmt"

func main() {
	fmt.Println("Como podemos criar um novo ponteiro?")

	xPtr := new(int)
	fmt.Print(xPtr)
}
