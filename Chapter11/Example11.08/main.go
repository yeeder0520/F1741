package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := []byte(`{"checkNum":123,"amount":200,"category":["gift","clothing"]}`)
	var v map[string]interface{}

	json.Unmarshal(jsonData, &v)
	fmt.Println(v)
	for key, value := range v {
		fmt.Println(key, "=", value)
	}
}
