package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format(time.ANSIC))
	fmt.Println(time.Now().Format(time.UnixDate))
	fmt.Println(time.Now().Format(time.RFC3339))
	fmt.Println(time.Now().Format("2006/1/2 3:4:5"))
}
