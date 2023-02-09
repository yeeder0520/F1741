package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	defer cleanUp()
	fmt.Println("程式執行中 (按下 Ctrl + C 來中斷)")

MainLoop:
	for {
		s := <-sigs
		switch s {
		case syscall.SIGINT:
			fmt.Println("程序中斷:", s)
			break MainLoop
		case syscall.SIGTERM:
			fmt.Println("程序終止:", s)
			break MainLoop
		}
	}

	fmt.Println("程式結果")
}

func cleanUp() {
	fmt.Println("進行清理作業...")
	for i := 0; i <= 10; i++ {
		fmt.Printf("刪除檔案 %v...(僅模擬)\n", i)
		time.Sleep(time.Millisecond * 100)
	}
}
