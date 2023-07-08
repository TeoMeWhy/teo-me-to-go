package main

import "os"

func main() {

	file, err := os.Create("test.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("Isso é um teste de escrita em um novo arquivo.")

}
