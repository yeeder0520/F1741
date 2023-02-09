package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func main() {
	p, err := loadPerson("data.json")
	if err != nil {
		fmt.Printf("%v", err)
		fmt.Printf("%T", err)
	}
	fmt.Println(p)
}

func loadPerson(fname string) (Person, error) {
	var p Person
	f, err := os.Open(fname)
	if err != nil {
		return p, err
	}
	err = json.NewDecoder(f).Decode(&p)
	if err != nil {
		return p, err
	}
	return p, err
}
