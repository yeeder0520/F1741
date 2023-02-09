package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 5
	b := &a

	v := reflect.ValueOf(b).Elem()
	fmt.Printf("%v %T\n", v.Int(), v.Int())

	v.SetInt(10)
	fmt.Printf("%v %T\n", v.Int(), v.Int())
	fmt.Printf("%v %T\n", v.Interface(), v.Interface())
	fmt.Printf("%v %T\n", *b, *b)
}
