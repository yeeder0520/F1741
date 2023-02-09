package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, out chan int, done chan bool) {
	for {
		n := rand.Intn(100)
		select {
		case out <- n:
			fmt.Printf("ID %d 傳送 %d\n", id, n)
		case <-done:
			fmt.Printf("ID %d 結束\n", id)
			return
		}
		time.Sleep(time.Millisecond)
	}
}

func work(workers, from, to int) int {
	out := make(chan int, workers)
	done := make(chan bool)
	defer close(done)

	for i := 0; i < workers; i++ {
		go worker(i, out, done)
	}

	res := 0
	for i := range out {
		if i >= 90 {
			res = i
			break
		}
	}
	return res
}

func main() {
	rand.Seed(time.Now().UnixNano())
	res := work(4, 1, 100)
	fmt.Println("答案:", res)
	time.Sleep(time.Second)
}
