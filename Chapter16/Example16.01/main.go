package main

import (
	"log"
	"sync"
)

func sum(from, to int, wg *sync.WaitGroup, mtx *sync.Mutex, res *int32) {
	for i := from; i <= to; i++ {
		mtx.Lock()
		*res += int32(i)
		mtx.Unlock()
	}
	wg.Done()
}

func main() {
	s1 := int32(0)
	wg := &sync.WaitGroup{}
	mtx := &sync.Mutex{}

	wg.Add(4)
	go sum(1, 25, wg, mtx, &s1)
	go sum(26, 50, wg, mtx, &s1)
	go sum(51, 75, wg, mtx, &s1)
	go sum(76, 100, wg, mtx, &s1)
	wg.Wait()

	log.SetFlags(0)
	log.Println(s1)
}
