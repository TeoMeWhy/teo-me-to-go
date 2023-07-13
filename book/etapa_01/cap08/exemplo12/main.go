package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}

	defer dir.Close()

	fileInfo, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	for _, fi := range fileInfo {
		fmt.Println(fi.Name())
	}
}
