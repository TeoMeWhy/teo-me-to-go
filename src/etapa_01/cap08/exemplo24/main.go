package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	maxp := flag.Int("max", 6, "The max value")
	flag.Parse()

	fmt.Println(rand.Intn(*maxp))
}
