package main

import (
	"fmt"
)

func greeting(fname, lname string) string {
	return fmt.Sprintln("哈囉:", fname, lname)
}

func main() {
	fname := "Edward"
	lname := "Scissorhands"
	fmt.Println("哈囉:", fname, lname)
	fmt.Printf("哈囉: %v %v\n", fname, lname)
	fmt.Print(greeting(fname, lname))
}
