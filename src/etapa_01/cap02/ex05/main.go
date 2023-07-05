package main

import "fmt"

func main() {
	fmt.Println(`Qual é o valor da expressão (true && false) || (false && true) || !(false && false)?`)

	fmt.Println("(true && false) || (false && true) || !(false && false) =", (true && false) || (false && true) || !(false && false))
}
