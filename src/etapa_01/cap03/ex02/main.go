package main

import "fmt"

func main() {
	fmt.Println("Qual é o valor de `x` após a execução de `x:=5; x+=1`?")
	x := 5
	x += 1
	fmt.Println(x)
}
