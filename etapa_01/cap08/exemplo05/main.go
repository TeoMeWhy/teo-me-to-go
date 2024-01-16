package main

import (
	"fmt"
	"strings"
)

func main() {
	nomes := []string{"Teo", "Maria", "Nah"}
	fmt.Println(strings.Join(nomes, "+"))
}
