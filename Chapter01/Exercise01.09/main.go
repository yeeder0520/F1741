package main

import "fmt"

func main() {
	var total float64 = 2 * 13
	fmt.Println("+ 主餐:", total)

	total = total + (4 * 2.25)
	fmt.Println("+ 飲料:", total)

	total = total - 5
	fmt.Println("- 折抵:", total)

	tip := total * 0.1
	total = total + tip
	fmt.Println("+ 小費:", total)

	split := total / 2
	fmt.Println("分攤額:", split)

	visitCount := 24
	visitCount = visitCount + 1
	remainder := visitCount % 5
	if remainder == 0 {
		fmt.Println("您已獲得來店滿 5 次折價券!")
	}
}
