package main

import "fmt"

func main() {

	var i int
	fmt.Scanf("%d", &i)

	switch i {

	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("três")
	case 4:
		fmt.Println("quatro")
	case 5:
		fmt.Println("cinco")
	default:
		fmt.Println("desconhecido")
	}

}
