package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 1; ; i++ {
		c <- fmt.Sprintf("%d: pinger", i)
	}
}

func ponger(c chan string) {
	for i := 1; ; i++ {
		c <- fmt.Sprintf("%d: ponger", i)
	}
}

func printer(c chan string) {
	for {
		fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
}

func main() {

	c := make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)

}
