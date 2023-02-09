package main

import "fmt"

func main() {
	{
		level := "Nest 1"
		fmt.Println("Block end   :", level)
	}
	// 將產生錯誤: undefined: level
	fmt.Println("Main end    :", level)
}
