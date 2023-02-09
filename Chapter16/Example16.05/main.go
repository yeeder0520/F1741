package main

import "fmt"

func doSomething() (chan int, chan bool) {
	in, out := make(chan int), make(chan bool)
	go func() {
		for i := range in {
			fmt.Println(i)
		}
		out <- true
	}()
	return in, out
}

func main() {
	in, out := doSomething()
	in <- 1
	in <- 2
	in <- 3
	close(in)
	<-out
}
