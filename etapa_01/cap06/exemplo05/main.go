package main

import "fmt"

func main() {
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(1, 1))
}
