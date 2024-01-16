package main

import "fmt"

func main() {
	slice := make([]int, 3, 9)
	fmt.Println("O tamanho da fatia é:", len(slice))
}
