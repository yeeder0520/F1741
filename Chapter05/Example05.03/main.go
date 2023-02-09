package main

import "fmt"

func main() {
	add := calculator("+")
	subtract := calculator("-")

	fmt.Println(add(5, 6))
	fmt.Println(subtract(10, 5))

	fmt.Printf("add()      型別: %T\n", add)
	fmt.Printf("subtract() 型別: %T\n", subtract)
}

func calculator(operator string) func(int, int) int {
	switch operator {
	case "+":
		return func(i, j int) int {
			return i + j
		}
	case "-":
		return func(i, j int) int {
			return i - j
		}
	}
	return nil
}
