package main

import (
	"fmt"
	"sync"
)

type Workers struct {
	in, out   chan int
	workerNum int
	mtx       sync.Mutex
}

func (w *Workers) init(maxWorkers, maxData int) {
	w.in, w.out = make(chan int, maxData), make(chan int)
	w.mtx = sync.Mutex{}
	for i := 0; i < maxWorkers; i++ {
		w.mtx.Lock()
		w.workerNum++
		w.mtx.Unlock()
		go w.readThem()
	}
}

func (w *Workers) addData(data int) {
	w.in <- data
}

func (w *Workers) readThem() {
	sum := 0
	for i := range w.in {
		sum += i
	}
	w.out <- sum

	w.mtx.Lock()
	w.workerNum--
	w.mtx.Unlock()
	if w.workerNum <= 0 {
		close(w.out)
	}
}

func (w *Workers) gatherResult() int {
	close(w.in)
	total := 0
	for i := range w.out {
		total += i
	}
	return total
}

func main() {
	maxWorkers := 10
	maxData := 100
	workers := Workers{}
	workers.init(maxWorkers, maxData)

	for i := 1; i <= maxData; i++ {
		workers.addData(i)
	}
	res := workers.gatherResult()
	fmt.Println(res)
}
