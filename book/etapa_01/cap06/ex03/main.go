package main

import (
	"errors"
	"fmt"
)

func maximo(x ...int) (int, error) {

	if len(x) < 1 {
		return 0, errors.New("Fodase")
	}

	max := x[0]
	for _, v := range x[1:] {
		if max < v {
			max = v
		}
	}
	return max, nil
}

func main() {
	fmt.Println("Escreva ua função com um parâmetro variádico que descubra o maior valor de uma lista de números.")

	x := []int{
		1, 24, 1,
		421, 51, 531,
		54, 2, 1, 0,
		12, 3, 451,
		51, 231, 938,
		24, 932,
	}

	fmt.Println(maximo(x...))
	fmt.Println(maximo(1, 3, 2103, 2412, 421, 98))
	fmt.Println(maximo(1))
	fmt.Println(maximo())
}
