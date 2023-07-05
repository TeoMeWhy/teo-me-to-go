package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println("A sequencia Fibonacci é definida como: fib(0)=0, fib(1)=1.")
	fmt.Println("fib(n)=fib(n-1) + fib(n-2). Escreva um programa com função recursiva capaz de falcular fib(n)")

	n := 0
	fmt.Scanf("%d", &n)

	fmt.Println(fib(n))
}
