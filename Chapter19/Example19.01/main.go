package main

import (
	"fmt"
	"reflect"
)

func Print(i interface{}) {
	fmt.Println("Type :", reflect.TypeOf(i))
	fmt.Println("Value:", reflect.ValueOf(i))
}

func main() {
	a := 5
	Print(a)
	b := &a
	Print(b)
	c := []string{"test"}
	Print(c)
	d := map[string]string{"a": "b"}
	Print(d)
}
