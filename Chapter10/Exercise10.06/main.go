package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	duration1 := time.Duration(time.Second * 360)
	duration2 := time.Duration(time.Hour*1 + time.Minute*30)
	fmt.Println("Dur1 :", duration1.Nanoseconds(), "ns")
	fmt.Println("DUr2 :", duration2.Nanoseconds(), "ns")

	date1 := now.Add(duration1)
	date2 := now.Add(duration2)
	fmt.Println("Now  :", now)
	fmt.Println("Date1:", date1)
	fmt.Println("Date2:", date2)
}
