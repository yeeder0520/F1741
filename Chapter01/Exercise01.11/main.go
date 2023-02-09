package main

import "fmt"

func main() {
	visits := 15

	fmt.Println("新顧客  :", visits == 1)
	fmt.Println("熟客    :", visits != 1)
	fmt.Println("銀牌會員:", visits > 10 && visits <= 20)
	fmt.Println("金牌會員:", visits > 20 && visits <= 30)
	fmt.Println("白金 VIP:", visits > 30)
}
