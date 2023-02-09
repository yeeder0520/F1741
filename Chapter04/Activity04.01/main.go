package main

import "fmt"

func getArr() [10]int {
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i + 10
	}
	return arr
}

func main() {
	fmt.Println(getArr())
}
