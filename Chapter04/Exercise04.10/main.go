package main

import "fmt"

func message() string {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := fmt.Sprintln("第一個元素:", s[0], s[0:1], s[:1])
	m += fmt.Sprintln("最末的元素:", s[len(s)-1], s[len(s)-1:len(s)], s[len(s)-1:])
	m += fmt.Sprintln("前五個元素:", s[:5])
	m += fmt.Sprintln("末四個元素:", s[5:])
	m += fmt.Sprintln("中間五元素:", s[2:7])
	m += fmt.Sprintln("全部的元素:", s[:])
	return m
}

func main() {
	fmt.Print(message())
}
