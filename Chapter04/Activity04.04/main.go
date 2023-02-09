package main

import (
	"fmt"
	"os"
)

var users = map[string]string{
	"305": "Sue",
	"204": "Bob",
	"631": "Jake",
	"073": "Tracy",
}

func getName(id string) (string, bool) {
	name, exists := users[id]
	return name, exists
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("未傳入使用者 ID")
		os.Exit(1)
	}
	for i := 1; i < len(os.Args); i++ {
		name, exists := getName(os.Args[i])
		if !exists {
			fmt.Printf("查無使用者 (%v)\n", os.Args[i])
		} else {
			fmt.Printf("哈囉, %v (%v)\n", name, os.Args[i])
		}
	}
}
