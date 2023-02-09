package main

import (
	"fmt"
	"time"
)

func main() {
	date1 := time.Date(2021, 4, 22, 16, 44, 05, 324359102, time.UTC)
	fmt.Println(date1)
	date2 := time.Date(2021, 4, 22, 16, 44, 05, 324359102, time.Local)
	fmt.Println(date2)
	date3 := date2.AddDate(-1, 3, 5)
	fmt.Println(date3)
}
