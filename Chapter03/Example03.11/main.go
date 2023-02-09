package main

import "fmt"

func main() {
	var message *string

	if message == nil {
		fmt.Println("錯誤, 非預期的 nil 值")
		return
	}
	fmt.Println(&message)
}
