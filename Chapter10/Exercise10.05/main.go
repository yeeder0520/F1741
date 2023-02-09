package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Date(2050, 12, 31, 0, 0, 0, 0, time.Local)

	fmt.Println("Equal :", time.Now().Equal(date))
	fmt.Println("Before:", time.Now().Before(date))
	fmt.Println("After :", time.Now().After(date))
}
