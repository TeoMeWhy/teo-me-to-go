package main

import "fmt"

func main() {
	fmt.Println("Dado o array a seguir, o que x[2:5] devolveria?")
	fmt.Println(`x := [6]string{"a", "b", "c", "d", "e", "f"}`)

	x := [6]string{"a", "b", "c", "d", "e", "f"}

	fmt.Println("x[2:5] =", x[2:5])
}
