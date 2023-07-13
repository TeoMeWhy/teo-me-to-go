package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 1; i < 10; i++ {

		duracao := time.Duration(rand.Intn(250))
		time.Sleep(duracao * time.Millisecond)
		fmt.Println(n, ":", i)
	}
}

func main() {

	for i := 0; i < 10; i++ {
		go f(i)
	}

	var input string
	fmt.Scanln(&input)
}
