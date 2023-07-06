package main

import "fmt"

func main() {
	fmt.Println("O que são `defer`, `panic` e `recover`? COmo podemos nos recuperar de uma pânico em um tempo de execução?")

	fmt.Println("São instruções que ajudam a mudar o fluxo de nossos programas.")

	fmt.Println("defer: utilizado para que uma função (ou comando) seja invocado ao final do escopo que pertence.")

	fmt.Println("panic: utilizado para lançar um erro durante a execução.")

	fmt.Println("recover: utilizado para capturr um erro durante a execução.")

	fmt.Println("Utilizamos o defer combinado com o recover para recuperar um pânico. Lembrando que isso é um `anti-padrão`")

}
