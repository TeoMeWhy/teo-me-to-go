package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	h := sha1.New()

	h.Write([]byte("teodoro"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}
