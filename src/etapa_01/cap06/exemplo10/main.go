package main

import "fmt"

func main() {

	defer func() {
		str := recover()
		fmt.Println("Capturado!")
		fmt.Println(str)

	}()

	panic("PANIC! ERRO BRUTAL!")
}
