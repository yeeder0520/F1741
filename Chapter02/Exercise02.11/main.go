package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		r := rand.Intn(8)
		if r%3 == 0 {
			fmt.Println("略過")
			continue
		} else if r%2 == 0 {
			fmt.Println("跳出")
			break
		}
		fmt.Println(r)
	}
}
