package main

import (
	"errors"
	"fmt"
)

func main() {
	defer fmt.Println("在 main() 使用 defer")
	test()
	fmt.Println("這一行不會印出")
}

func test() {
	defer fmt.Println("在 test() 使用 defer")
	msg := "good-bye"
	message(msg)
}

func message(msg string) {
	defer fmt.Println("在 message() 使用 defer")
	if msg == "good-bye" {
		panic(errors.New("出事了"))
	}
}
