package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("我是第 1 個宣告的!")
	}()
	defer func() {
		fmt.Println("我是第 2 個宣告的!")
	}()
	defer func() {
		fmt.Println("我是第 3 個宣告的!")
	}()
	f1 := func() {
		fmt.Println("f1 開始")
	}
	f2 := func() {
		fmt.Println("f2 結束")
	}

	f1()
	f2()
	fmt.Println("main() 結束")
}
