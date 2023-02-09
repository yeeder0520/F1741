package main

import (
	"fmt"
	"time"
)

func main() {
	t1, err := time.Parse(time.ANSIC, "Thu Apr 22 16:44:05 2021")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("from ANSIC   :", t1)

	t2, err := time.Parse(time.UnixDate, "Thu Apr 22 16:44:05 CST 2021")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("from UnixDate:", t2)

	t3, err := time.Parse(time.RFC3339, "2021-04-22T16:44:05+08:00")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("from RFC3339 :", t3)

	t4, err := time.Parse("2006/1/2 3:4:5", "2021/4/22 4:44:5")
	// t4, err := time.Parse("2006/1/2 PM 3:4:5", "2021/4/22 PM 4:44:5")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("from custom  :", t4)
}
