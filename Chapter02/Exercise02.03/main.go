package main

import (
	"fmt"
)

func main() {
	input := 10

	if input < 0 {
		fmt.Println("輸入值不得為負")
	} else if input%2 == 0 {
		fmt.Println(input, "是偶數")
	} else {
		fmt.Println(input, "是奇數")
	}
}
