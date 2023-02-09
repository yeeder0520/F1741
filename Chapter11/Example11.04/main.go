package main

import (
	"encoding/json"
	"fmt"
)

type greeting struct {
	SomeMessage string
}

func main() {
	var v greeting
	v.SomeMessage = "Marshal me!"

	json, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", json)
}
