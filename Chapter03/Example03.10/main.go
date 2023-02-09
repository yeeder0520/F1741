package main

import "fmt"

func main() {
	username := "Sir_King_Über"

	fmt.Println("Bytes :", len(username))
	fmt.Println("Runes :", len([]rune(username)))
	fmt.Println(string(username[:10]))
	fmt.Println(string([]rune(username)[:10]))
}
