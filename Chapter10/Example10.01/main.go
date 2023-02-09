package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("程式開始時間:", start)
	fmt.Println("資料處理中...")
	time.Sleep(2 * time.Second)
	end := time.Now()
	fmt.Println("程式結束時間:", end)
}
