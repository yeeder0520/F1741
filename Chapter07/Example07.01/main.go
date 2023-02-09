package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	s := `{"Name":"Joe","Age":18}`
	s2 := `{"Name":"Jane","Age":21}`

	p, err := loadPerson(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	p2, err := loadPerson2(strings.NewReader(s2))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)

	f, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	p3, err := loadPerson2(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p3)
}

func loadPerson(s string) (Person, error) {
	var p Person
	err := json.NewDecoder(strings.NewReader(s)).Decode(&p)
	if err != nil {
		return p, err
	}
	return p, nil
}

func loadPerson2(r io.Reader) (Person, error) {
	var p Person
	err := json.NewDecoder(r).Decode(&p)
	if err != nil {
		return p, err
	}
	return p, err
}
