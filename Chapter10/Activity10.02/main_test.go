package main

import (
	"log"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	for i := 0; i < 1000; i++ {
		start := time.Now()
		if Fibonacci(30) != 832040 {
			t.Fatal("測試失敗")
		}
		log.Println(time.Since(start).Seconds())
	}
}
