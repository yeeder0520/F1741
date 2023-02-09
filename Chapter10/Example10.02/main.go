package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	day := now.Weekday()
	// day = time.Monday
	hour := now.Hour()
	// hour = 1
	fmt.Println("Day:", day, "/ hour:", hour)
	if day.String() == "Monday" && (hour >= 0 && hour < 2) {
		fmt.Println("執行全功能測試")
	} else {
		fmt.Println("執行簡易測試")
	}
}
