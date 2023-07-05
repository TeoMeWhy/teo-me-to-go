package main

import "fmt"

func zero(x int) {
	x = 0
}

func main() {

	x := 10
	fmt.Println(x)

	zero(x)

	fmt.Println(x)

}
