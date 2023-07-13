package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Mensagem de Erro!")
	fmt.Println(err)
}
