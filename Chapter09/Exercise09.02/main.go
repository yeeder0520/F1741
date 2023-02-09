package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 255; i++ {
		fmt.Printf("10 進位: %3.d | ", i)
		fmt.Printf("2 進位: %8.b | ", i)
		fmt.Printf("16 進位: %2.x\n", i)
	}
}
