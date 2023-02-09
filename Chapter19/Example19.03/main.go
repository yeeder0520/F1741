package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name    string  `des:"userName"`
	Age     int     `des:"userAge"`
	Balance float64 `des:"bankBalance"`
	Member  bool    `des:"isMember"`
}

func PrintStruct(s interface{}) {
	sT := reflect.TypeOf(s)
	sV := reflect.ValueOf(s)

	fmt.Printf("type %s %v {\n", sT.Name(), sT.Kind().String())

	for i := 0; i < sT.NumField(); i++ {
		field := sT.Field(i)
		value := sV.Field(i)

		fmt.Printf("\t%s\t%s\t= %v\t(description: %s)\n",
			field.Name, field.Type.String(),
			value.Interface(), field.Tag.Get("des"))
	}

	fmt.Println("}")
}

func main() {
	u1 := User{
		Name:    "Tracy",
		Age:     51,
		Balance: 98.43,
		Member:  true,
	}

	PrintStruct(u1)

	v := reflect.ValueOf(&u1)
	v.Elem().FieldByName("Name").SetString("Grace")
	v.Elem().FieldByName("Age").SetInt(45)
	v.Elem().FieldByName("Balance").SetFloat(56.97)
	v.Elem().FieldByName("Member").SetBool(false)

	PrintStruct(u1)
}
