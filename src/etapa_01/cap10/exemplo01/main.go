package main

import (
	"fmt"
)

func f(n int) {
	for i := 1; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {

	go f(0)
	var input string
	fmt.Scanln(&input)

}
