package main

import (
	"fmt"
)

func main() {
	x := 9
	sqr := func(i int) int {
		return i * i
	}
	fmt.Printf("%d 的平方為 %d\n", x, sqr(x))
}
