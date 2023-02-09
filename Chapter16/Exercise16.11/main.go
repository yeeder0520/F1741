package main

import (
	"context"
	"fmt"
	"time"
)

func countNumbers(c context.Context, out chan int) {
	i := 0
	for {
		select {
		case <-c.Done():
			out <- i
			return
		default:
			time.Sleep(time.Millisecond * 100)
			i++
		}
	}
}

func main() {
	out := make(chan int)

	c := context.TODO()
	cl, cancel := context.WithCancel(c)

	go countNumbers(cl, out)

	time.Sleep(time.Millisecond * 100 * 5)
	cancel()

	fmt.Println(<-out)
}
