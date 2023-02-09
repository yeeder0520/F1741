package main

import (
	"encoding/json"
	"fmt"
)

type book struct {
	ISBN          string `json:"isbn"`
	Title         string `json:"title"`
	YearPublished int    `json:"yearpub,omitempty"`
	Author        string `json:"author"`
	CoAuthor      string `json:"coauthor,omitempty"`
}

/*
type book struct {
	ISBN          string `json:"-"`
	Title         string `json:"title"`
	YearPublished int    `json:"yearpub,omitempty"`
	Author        string `json:""`
	CoAuthor      string `json:",omitempty"`
}
*/

func main() {
	b := book{}
	b.ISBN = "9933HIST"
	b.Title = "Greatest of all Books"
	// b.YearPublished = 2020
	b.Author = "John Adams"

	json, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))
}
