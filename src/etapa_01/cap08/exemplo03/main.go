package main

import (
	"fmt"
	"strings"
)

func main() {
	nome := "Téo Calvo"
	fmt.Println(strings.HasPrefix(nome, "Téo"))
	fmt.Println(strings.HasSuffix(nome, "Calvo"))
}
