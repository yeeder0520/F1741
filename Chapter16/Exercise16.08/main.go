package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(in, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for i := range in {
		sum += i
		time.Sleep(time.Millisecond)
	}
	out <- sum
}

func sum(in, out chan int) {
	sum := 0
	for i := range in {
		sum += i
	}
	out <- sum
}

func work(workers, from, to int) int {
	wg := &sync.WaitGroup{}
	wg.Add(workers)

	in := make(chan int, (to-from)+1)
	out := make(chan int, workers)
	res := make(chan int, 1)

	for i := 0; i < workers; i++ {
		go worker(in, out, wg)
	}
	go sum(out, res)

	for i := from; i <= to; i++ {
		in <- i
	}
	close(in)
	wg.Wait()
	close(out)

	return <-res
}

func main() {
	res := work(4, 1, 100)
	fmt.Println(res)
}
