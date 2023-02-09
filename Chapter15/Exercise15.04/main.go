package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templateStr = `
<html>
  <h1>Customer {{.ID}}</h1>
  {{if .ID }}
    <p>Details:</p>
    <ul>
        {{if .Name}}<li>Name: {{.Name}}</li>{{end}}
        {{if .Surname}}<li>Surname: {{.Surname}}</li>{{end}}
        {{if .Age}}<li>Age: {{.Age}}</li>{{end}}
    </ul>
  {{else}}
    <p>Data not available</p>
  {{end}}
</html>
`

type Customer struct {
	ID      int
	Name    string
	Surname string
	Age     int
}

func hello(w http.ResponseWriter, r *http.Request) {
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

	tmpl, _ := template.New("Exercise15.04").Parse(templateStr)
	tmpl.Execute(w, customer)
}

func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
