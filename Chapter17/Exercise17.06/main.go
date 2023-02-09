// package main: Exercise 17.06
package main

import "fmt"

// Calc defines a calculator construct
type Calc struct{}

// Add returns the total of two integers added together
func (c Calc) Add(a, b int) int {
	return a + b
}

// Multiply returns the total of one integers multipled by the other
func (c Calc) Multiply(a, b int) int {
	return a * b
}

// PrintResult prints out the received integer argument
func PrintResult(i int) {
	fmt.Println(i)
}

func main() {
	calc := Calc{}
	PrintResult(calc.Add(1, 1))
	PrintResult(calc.Multiply(2, 2))
}
