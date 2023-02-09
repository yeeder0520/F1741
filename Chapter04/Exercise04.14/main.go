package main

import (
	"fmt"
	"os"
)

func getUsers() map[string]string {
	return map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}
}

func getUser(id string) (string, bool) {
	users := getUsers()
	user, exists := users[id]
	return user, exists
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("未傳入使用者 ID")
		os.Exit(1)
	}
	userID := os.Args[1]
	name, exists := getUser(userID)
	if !exists {
		fmt.Printf("查無傳入的使用者 ID (%v).\n使用者列表:\n", userID)
		for key, value := range getUsers() {
			fmt.Println("使用者 ID:", key, "  名字:", value)
		}
		os.Exit(1)
	}
	fmt.Println("查得名字:", name)
}
