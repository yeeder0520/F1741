package main

import "fmt"

// Fibonacci calculates a value in the Fibonacci sequence
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	fmt.Println(Fibonacci(30))
}
