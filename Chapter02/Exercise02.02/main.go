package main

import (
	"fmt"
)

func main() {
	input := 4

	if input%2 == 0 {
		fmt.Println(input, "是偶數")
	} else {
		fmt.Println(input, "是奇數")
	}
}
