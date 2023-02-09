package main

import (
	"fmt"
	"sync"
)

func sum(from, to int, wg *sync.WaitGroup, res *int) {
	*res = 0
	for i := from; i <= to; i++ {
		*res += i
	}
	if wg != nil {
		wg.Done()
	}
}

func main() {
	var s1, s2 int

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go sum(1, 100, wg, &s1)
	sum(1, 10, nil, &s2)
	wg.Wait()

	fmt.Println(s1, s2)
}
