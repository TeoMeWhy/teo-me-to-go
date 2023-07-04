package main

import "fmt"

func main() {
	txt := `	var x []float64

	y := make([]float64, 5)

	z := make([]float64, 5, 10)

	arr := [5]float64{1, 2, 3, 4, 5}
	t := arr[0:5]`

	fmt.Println(txt)
}
