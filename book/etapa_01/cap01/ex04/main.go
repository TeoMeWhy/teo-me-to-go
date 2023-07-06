package main

import "fmt"

func main() {
	fmt.Println("Usamos a função `Println` definida no pacote `fmt`. Se quiséssemos usar a função `Exit` do pacote `os`, o que você deveria fazer?")

	fmt.Println("RESPOSTA: Primeiro devemos importar o pacote `os` com o comando `import os`. Depois invocar a função `Exit` utilizando `os.Exit()`")

}
