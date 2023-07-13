package main

import (
	"fmt"
	"strings"
)

func main() {
	nome := "Teo Me Why"

	index := strings.Index(nome, "M")
	fmt.Println(index)
	fmt.Println(string(nome[index]))
}
