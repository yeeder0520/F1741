package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	LastName  string  `json:"lname"`
	FirstName string  `json:"fname"`
	Address   address `json:"address"`
}

type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}

func main() {
	addr := address{
		Street:  "Galaxy Far Away",
		City:    "Dark Side",
		State:   "Tatooine",
		ZipCode: 12345,
	}
	p := person{
		LastName:  "Vader",
		FirstName: "Darth",
		Address:   addr,
	}

	noPrettyPrint, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(noPrettyPrint))
	fmt.Println()

	prettyPrint, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(prettyPrint))
}
