package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Customer struct {
	ID      int
	Name    string
	Surname string
	Age     int
}

type Hello struct {
	tmpl *template.Template
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vl := r.URL.Query()
	customer := Customer{}

	id, ok := vl["id"]
	if ok {
		customer.ID, _ = strconv.Atoi(id[0])
	}

	name, ok := vl["name"]
	if ok {
		customer.Name = name[0]
	}

	surname, ok := vl["surname"]
	if ok {
		customer.Surname = surname[0]
	}

	age, ok := vl["age"]
	if ok {
		customer.Age, _ = strconv.Atoi(age[0])
	}

	h.tmpl.Execute(w, customer)
}

func main() {
	hello := Hello{}
	hello.tmpl, _ = template.ParseFiles("./template/hello_tmpl.html")
	http.Handle("/", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
