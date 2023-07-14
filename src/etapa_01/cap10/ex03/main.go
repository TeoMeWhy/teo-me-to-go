package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("O que é um canal com buffer? Como você criaria um com capacidade igual a 20?")

	ch := make(chan int, 20)

	go func() {

		for {
			ch <- rand.Intn(10)
			fmt.Println("Enviando dados...")
		}
	}()

	go func() {
		for {
			msg := <-ch
			fmt.Println(msg)
			time.Sleep(time.Second * 2)
		}
	}()

	var input string
	fmt.Scanln(&input)

}
