package main

import "fmt"

func main() {
	fmt.Print("Entre com um valor: ")

	var a int
	fmt.Scanf("%d", &a)

	res := a * 2
	fmt.Println(res)
}
