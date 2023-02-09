package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	v := make(map[string]interface{})
	v["checkNum"] = 123
	v["amount"] = 200
	v["category"] = []string{"gift", "clothing"}

	jsonData, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonData))
}
