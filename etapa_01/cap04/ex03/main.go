package main

import "fmt"

func main() {
	fmt.Println("Escreva um programa que exiba números de 1 a 100, mas para múltiplos de 3, mostre `Fizz` no lugar do número, e para mútiplos de cinco, exiba `Buzz`. Para números que seja, múltiplos tanto de 3 quanto de 5, mostre `FizzBuzz`.")

	for i := 1; i <= 100; i++ {
		txt := ""

		if i%3 == 0 {
			txt += "Fizz"
		}

		if i%5 == 0 {
			txt += "Buzz"
		}

		if txt == "" {
			fmt.Println(i)
		} else {
			fmt.Println(txt)
		}
	}
}
