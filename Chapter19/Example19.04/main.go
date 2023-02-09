package main

import (
	"fmt"
	"reflect"
)

func runDeepEqual(a, b interface{}) {
	fmt.Printf("%v DeepEqual %v : %v\n", a, b, reflect.DeepEqual(a, b))
}

func main() {
	runDeepEqual([3]int{1, 2, 3}, [3]int{1, 2, 3})
	runDeepEqual([]int{1, 2, 3}, []int{1, 2, 3})

	a := map[int]string{1: "one", 2: "two"}
	b := map[int]string{1: "one", 2: "two"}
	runDeepEqual(a, b)

	var c, d interface{}
	c = map[int]string{1: "one", 2: "two"}
	d = map[int]string{1: "one", 2: "two"}
	runDeepEqual(c, d)
}
