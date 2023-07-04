package main

import "fmt"

func main() {
	elementos := make(map[string]string)
	elementos["H"] = "Hidrogênio"
	elementos["He"] = "Hélio"
	elementos["Li"] = "Lítio"
	elementos["Be"] = "Meryllium"
	elementos["B"] = "Bóro"
	elementos["C"] = "Carbono"
	elementos["N"] = "Nitrogênio"
	elementos["O"] = "Oxigênio"
	elementos["F"] = "Fluor"
	elementos["Ne"] = "Neon"

	fmt.Println(elementos["Li"])

	// name, ok := elementos["Teo"]
	// fmt.Println(name, ok)

	if name, ok := elementos["Teo"]; ok {
		fmt.Println(name)
	}
}
