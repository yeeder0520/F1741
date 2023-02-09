package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Sunday
	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("生日為平日")
	case time.Saturday, time.Sunday:
		fmt.Println("生日為周末")
	default:
		fmt.Println("生日錯誤")
	}
}
