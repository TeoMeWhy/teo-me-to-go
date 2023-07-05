package main

import "fmt"

func soma(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {

	x := []int{1, 4, 8, 10}
	fmt.Println(soma(x...))

	fmt.Println(soma(1, 2, 3))
}
