package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		txt := ""
		if i%3 == 0 {
			txt += "Fizz"
		}
		if i%5 == 0 {
			txt += "Buzz"
		}
		if txt == "" {
			fmt.Println(i)
		} else {
			fmt.Println(txt)
		}
	}

}
