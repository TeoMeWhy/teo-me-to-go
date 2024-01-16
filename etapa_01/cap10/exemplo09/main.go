package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		for {
			c1 <- "Canal 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "Canal 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {

		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second):
				fmt.Println("Timeout")
			default:
				fmt.Println("Nada a ser lido")
				time.Sleep(time.Millisecond * 500)
			}
		}

	}()

	var input string
	fmt.Scanln(&input)
}
