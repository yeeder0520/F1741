package main

import (
	"encoding/json"
	"fmt"
)

type greeting struct {
	Message string
}

func main() {
	data := []byte(` 
{ 
	"message": "Greetings fellow gopher!"
}
`)
	var v greeting
	err := json.Unmarshal(data, &v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}
