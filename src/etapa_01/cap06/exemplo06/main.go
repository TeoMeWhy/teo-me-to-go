package main

import "fmt"

func main() {
	x := 0

	incremento := func() int {
		x++
		return x
	}

	fmt.Println(incremento())
	fmt.Println(incremento())

	fmt.Println(x)
}
