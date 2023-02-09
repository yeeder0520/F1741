package main

import "fmt"

func greet(ch chan string) {
	ch <- "Hello"
}

func main() {
	ch := make(chan string)
	go greet(ch)
	fmt.Println(<-ch)
}
