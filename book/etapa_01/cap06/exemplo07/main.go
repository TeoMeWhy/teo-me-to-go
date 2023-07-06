package main

import "fmt"

func constroiGeradorPares() func() int {
	i := 0
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}

func main() {

	proxPar := constroiGeradorPares()
	fmt.Println(proxPar())
	fmt.Println(proxPar())
	fmt.Println(proxPar())

}
