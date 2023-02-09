package main

import (
	"errors"
	"fmt"
	"reflect"
)

func doubler(i interface{}) (string, error) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	switch t.String() {
	case "string":
		return v.String() + v.String(), nil
	case "bool":
		if v.Bool() {
			return "truetrue", nil
		}
		return "falsefalse", nil
	case "float32", "float64":
		if t.String() == "float64" {
			return fmt.Sprint(v.Float() * 2), nil
		}
		return fmt.Sprint(float32(v.Float()) * 2), nil
	case "int", "int8", "int16", "int32", "int64":
		return fmt.Sprint(v.Int() * 2), nil
	case "uint", "uint8", "uint16", "uint32", "uint64":
		return fmt.Sprint(v.Uint() * 2), nil
	default:
		return "", errors.New("傳入了未支援的值")
	}
}

func main() {
	res, _ := doubler(-5)
	fmt.Println("-5  :", res)
	res, _ = doubler(5)
	fmt.Println("5   :", res)
	res, _ = doubler("yum")
	fmt.Println("yum :", res)
	res, _ = doubler(true)
	fmt.Println("true:", res)
	res, _ = doubler(float32(3.14))
	fmt.Println("3.14:", res)
}
