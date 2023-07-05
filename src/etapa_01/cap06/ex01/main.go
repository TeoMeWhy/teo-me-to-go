package main

import "fmt"

func sum(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func main() {
	fmt.Println("`sum` é uma função que aceita uma fatia de números e os soma. Como seria a assinatura desta função em Go?")
	fmt.Println("func sum(x []int) int")
}
