package main

import "fmt"

func media(x []float64) (res float64) {
	total := 0.0
	for _, v := range x {
		total += v
	}

	res = total / float64(len(x))
	return
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(media(xs))
}
