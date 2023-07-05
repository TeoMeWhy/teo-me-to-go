package main

import "fmt"

func main() {
	elementos := map[string]map[string]string{
		"H": {
			"nome":   "Hidrogênio",
			"estado": "Gasoso",
		},

		"He": {
			"nome":   "Hélio",
			"estado": "Gasoso",
		},

		"Li": {
			"nome":   "Lítio",
			"estado": "Sólido",
		},

		"Be": {
			"nome":   "Beryllium",
			"estado": "Sólido",
		},

		"B": {
			"nome":   "Bóro",
			"estado": "Sólido",
		},

		"C": {
			"nome":   "Carbono",
			"estado": "Sólido",
		},

		"N": {
			"nome":   "Nitrogênio",
			"estado": "Gasoso",
		},

		"O": {
			"nome":   "Oxigênio",
			"estado": "Gasoso",
		},

		"F": {
			"nome":   "Fluor",
			"estado": "Gasoso",
		},

		"Ne": {
			"nome":   "Neon",
			"estado": "Gasoso",
		},
	}

	var elemento string

	fmt.Println("Entre com a sigla de um elemento: ")
	fmt.Scanf("%s", &elemento)

	if el, ok := elementos[elemento]; ok {
		fmt.Println(el["nome"], ":", el["estado"])
	}
}
