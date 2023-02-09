package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
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
	dataStr := string(data)
	p := person{}

	decoder := json.NewDecoder(strings.NewReader(dataStr))
	if err := decoder.Decode(&p); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(p)
	fmt.Println()

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(p); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
