package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Lastname  string  `json:"lname"`
	Firstname string  `json:"fname"`
	Address   address `json:"address"`
}

type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}

func main() {
	data := []byte(`
{
	"lname":"Smith",
	"fname":"John",
	"address":{
		"street":"Sulphur Springs Rd",
		"city":"Park City",
		"state":"VA",
		"zipcode":12345
	}
}	 
`)
	p := person{}
	if err := json.Unmarshal(data, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", p)
}
