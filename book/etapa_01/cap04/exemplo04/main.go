package main

import "fmt"

func main() {

	var check string

	for i := 1; i <= 10; i++ {

		if i%2 == 0 {
			check = "par"
		} else {
			check = "ímpar"
		}

		fmt.Println(i, check)
	}
}
