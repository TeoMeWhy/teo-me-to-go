package main

import "fmt"

func main() {
	fmt.Println("Por que usaríamos um campo anônimo incluído em vez de utilizar um campo normal nomeado?")

	fmt.Println("Pois o campo anônimo no dá a propriedade da struct `ser` daquele tipo, e não `ter` o tipo.")
	fmt.Println("Com isso, a struct nova tem acesso aos métodos do campo anônimo.")
}
