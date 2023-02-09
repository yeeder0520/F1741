package main

import "fmt"

func main() {
	fname := "Joe"
	gpa := 3.75
	hasJob := true
	age := 24
	hourlyWage := 45.53
	fmt.Printf("%s 的 GPA: %f\n", fname, gpa)
	fmt.Printf("有工作: %t\n", hasJob)
	fmt.Printf("年齡: %d, 時薪: %v\n", age, hourlyWage)
}
