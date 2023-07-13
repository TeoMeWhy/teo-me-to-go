package main

import (
	"fmt"
	"strings"
)

func main() {
	nome := "Teodoro"

	fmt.Println(strings.Replace(nome, "o", "a", 2))
}
