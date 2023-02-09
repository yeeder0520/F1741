package main

import (
	"errors"
	"fmt"
)

func main() {
	a()
	fmt.Println("這一行現在會印出了")
}

func a() {
	b("good-bye")
	fmt.Println("返回 a()")
}

func b(msg string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("b() 發生錯誤:", r)
		}
	}()
	if msg == "good-bye" {
		panic(errors.New("事情出代誌了"))
	}
	fmt.Print(msg)
}
