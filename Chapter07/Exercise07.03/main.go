package main

import (
	"fmt"
)

type person struct {
	lastName  string
	age       int
	isMarried bool
}

type animal struct {
	name     string
	category string
}

type record struct {
	key       string
	data      interface{}
	valueType string
}

func main() {
	m := make(map[string]interface{})

	m["person"] = person{lastName: "Doe", isMarried: false, age: 19}
	m["firstname"] = "Smith"
	m["age"] = 54
	m["isMarried"] = true
	m["animal"] = animal{name: "oreo", category: "cat"}

	rs := []record{}
	for k, v := range m {
		rs = append(rs, newRecord(k, v))
	}

	for _, v := range rs {
		fmt.Println("Key: ", v.key)
		fmt.Println("Data: ", v.data)
		fmt.Println("Type: ", v.valueType)
		fmt.Println()
	}
}

func newRecord(key string, i interface{}) record {
	r := record{}
	r.key = key
	switch v := i.(type) {
	case int:
		r.valueType = "int"
		r.data = v
		return r
	case bool:
		r.valueType = "bool"
		r.data = v
		return r
	case string:
		r.valueType = "string"
		r.data = v
		return r
	case person:
		r.valueType = "person"
		r.data = v
		return r
	default:
		r.valueType = "unknown"
		r.data = v
		return r
	}
}
