package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

type student struct {
	StudentId     int
	LastName      string
	MiddleInitial string
	FirstName     string
	IsEnrolled    bool
	Courses       []course
}

type course struct {
	Name   string
	Number int
	Hours  int
}

func main() {
	s := student{
		StudentId:  2,
		LastName:   "Washington",
		FirstName:  "Bill",
		IsEnrolled: true,
		Courses: []course{
			{Name: "World Lit", Number: 101, Hours: 3},
			{Name: "Biology", Number: 201, Hours: 4},
			{Name: "Intro to Go", Number: 101, Hours: 4},
		},
	}

	var conn bytes.Buffer
	encoder := gob.NewEncoder(&conn)
	if err := encoder.Encode(&s); err != nil {
		fmt.Println("GOB 編碼錯誤:", err)
		os.Exit(1)
	}

	fmt.Printf("%x\n", conn.String())

	s2 := student{}
	decoder := gob.NewDecoder(&conn)
	if err := decoder.Decode(&s2); err != nil {
		fmt.Println("GOB 解碼錯誤:", err)
		os.Exit(1)
	}

	fmt.Println(s2)
}
