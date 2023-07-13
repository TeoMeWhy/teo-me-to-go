package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func gethash(filename string) (uint32, error) {

	//abre arquivo
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	defer f.Close()

	h := crc32.NewIEEE()

	_, err = io.Copy(h, f)
	if err != nil {
		return 0, err
	}

	return h.Sum32(), nil
}

func main() {

	h1, err := gethash("teste1.txt")
	if err != nil {
		return
	}

	h2, err := gethash("teste2.txt")
	if err != nil {
		return
	}

	fmt.Println(h1, h2, h1 == h2)
}
