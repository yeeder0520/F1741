package main

import "fmt"

var budgetCategories = make(map[int]string)

func init() {
	fmt.Println("初始化預算分類...")
	budgetCategories[1] = "汽車保險"
	budgetCategories[2] = "貸款"
	budgetCategories[3] = "電費"
	budgetCategories[4] = "退休金"
	budgetCategories[5] = "旅遊補助"
	budgetCategories[7] = "雜貨支出"
	budgetCategories[8] = "汽車貸款"
}

func main() {
	for k, v := range budgetCategories {
		fmt.Printf("鍵: %d, 值: %s\n", k, v)
	}
}
