package main

import "fmt"

func first() {
	fmt.Println("Primeiro!")
}

func second() {
	fmt.Println("Segundo!")
}

func main() {

	defer second()
	first()
	fmt.Println("Ainda não acabou...")

}
