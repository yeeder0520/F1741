package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	appName := "HTTPCHECKER"
	action := "BASIC"
	date := time.Now()
	logFileName := appName + "_" + action + "_" + strconv.Itoa(date.Year()) + "_" + date.
		Month().String() + "_" + strconv.Itoa(date.Day()) + ".log"
	//logFileName := fmt.Sprintf("%v_%v_%v_%v_%v.log", appName, action, date.Year(), date.
	//	Month().String(), date.Day())
	fmt.Println("log 檔名稱:", logFileName)
}
