package main

import "fmt"

func main() {
	fmt.Print("Entre como um valor a ser convertido de Fahrenheit para Celsius: ")

	var f float64
	fmt.Scanf("%f", &f)

	c := (f - 32.) * (5. / 9.)
	fmt.Println(c)
}
