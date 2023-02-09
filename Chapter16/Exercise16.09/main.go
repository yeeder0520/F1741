package main

import (
	"fmt"
	"strings"
)

func readThem(in chan string, done chan bool) {
	for i := range in {
		fmt.Println(strings.ToUpper(i))
	}
	done <- true
}

func main() {
	strs := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	workers := 4
	in := make(chan string, len(strs))
	done := make(chan bool, workers)

	for i := 0; i < workers; i++ {
		go readThem(in, done)
	}

	for _, s := range strs {
		in <- s
	}
	close(in)

	for i := 0; i < workers; i++ {
		<-done
	}
}
