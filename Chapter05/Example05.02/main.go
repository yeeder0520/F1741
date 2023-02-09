package main

import (
	"fmt"
)

type calc func(int, int) int

func main() {
	calculator(add, 5, 6)
	calculator(subtract, 10, 5)
}

func add(i, j int) int {
	return i + j
}

func subtract(i, j int) int {
	return i - j
}

func calculator(f calc, i, j int) {
	fmt.Println(f(i, j))
}
