package main

import "fmt"

func main() {
	fmt.Println(`Sabemos que na base 10 o maior número de um dígito é 9 e o maior número de dois dígitos é 99.
	(...) qual é o maior número de oitro dígitos?`)

	fmt.Println("Base 10 = 10^8 - 1 =", 100000000-1)
	fmt.Println("Base 2 = 2^8 - 1 =", 2*2*2*2*2*2*2*2-1)
}
