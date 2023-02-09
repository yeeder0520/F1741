package main

import "fmt"

type cat struct {
	name string
}

func main() {
	i := 99
	b := false
	str := "test"
	c := cat{name: "oreo"}
	printDetails(i, b, str, c)
}

func printDetails(data ...interface{}) {
	for _, i := range data {
		fmt.Printf("%v, %T\n", i, i)
	}
}
