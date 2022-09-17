package main

import "fmt"

func main() {
	var x string = "hello"
	var y string = "world"

	fmt.Println(x == y)

	x = "hello"
	y = "hello"

	fmt.Println(x == y)

}
