package main

import "fmt"

func main() {

	fmt.Println("Ex 04: Usamos a função *Println* definida no pacote *fmt*. Se quiséssemos usar a função *Exit* do pacote *os*, o que você deveria fazer?")

	fmt.Println("\nResposta: Para usar a função *Exit* do pacote *os*, primeiro é necessário importar o pacote *os* com o comando *import \"os\"*, posteriormente, utilizar *os.Exit()* para invocar (executar) a função *Exit*.")
}
