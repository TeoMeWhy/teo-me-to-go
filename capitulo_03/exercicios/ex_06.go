package main

import "fmt"

func main() {
	fmt.Println("Escreva outro programa que faça a conversão de pés para metros (1 pé = 0,3048 m).")

	var pes float64
	fmt.Println("Entre com valor em pés para ser convertido em metros: ")
	fmt.Scanf("%f", &pes)

	metros := pes * 0.3048
	fmt.Println(pes, " pés =", metros, "metros")
}
