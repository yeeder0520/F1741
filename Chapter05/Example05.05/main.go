package main

import "fmt"

func main() {
	age := 25
	name := "John"
	defer personAge(name, age)

	age *= 2
	fmt.Println("年齡加倍:")
	personAge(name, age)
}

func personAge(name string, i int) {
	fmt.Printf("%s 是 %d 歲\n", name, i)
}
