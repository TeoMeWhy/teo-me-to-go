package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 3)
	c2 := make(chan string, 1)

	go func() {
		for i := 1; ; i++ {
			fmt.Println("Adicionando dados...")
			c1 <- fmt.Sprintf("Canal 1: %d", i)
		}
	}()

	go func() {
		for i := 1; ; i++ {
			c2 <- fmt.Sprintf("Canal 2: %d", i)
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
