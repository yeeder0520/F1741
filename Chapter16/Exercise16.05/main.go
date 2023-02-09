package main

import "fmt"

func greet(ch chan string) {
	msg := <-ch
	ch <- fmt.Sprintf("收到訊息: %s", msg)
	ch <- "Hello David"
}

func main() {
	ch := make(chan string)
	go greet(ch)

	ch <- "Hello John"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
