package main

import "fmt"

func main() {
	elementos := map[string]string{
		"H":  "Hidrogênio",
		"He": "Hélio",
		"Li": "Lítio",
		"Be": "Meryllium",
		"B":  "Bóro",
		"C":  "Carbono",
		"N":  "Nitrogênio",
		"O":  "Oxigênio",
		"F":  "Fluor",
		"Ne": "Neon",
	}

	fmt.Println(elementos["Li"])

	// name, ok := elementos["Teo"]
	// fmt.Println(name, ok)

	if name, ok := elementos["Teo"]; ok {
		fmt.Println(name)
	}
}
