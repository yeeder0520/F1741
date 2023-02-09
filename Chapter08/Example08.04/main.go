package main

import (
	"fmt"
)

var name = "Gopher"

func init() {
	fmt.Println("哈囉,", name)
}

func main() {
	fmt.Println("哈囉, main() 函式")
}
