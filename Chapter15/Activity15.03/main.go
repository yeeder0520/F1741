package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Visitor struct {
	Name string
}

type Hello struct {
	tmpl *template.Template
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vl := r.URL.Query()

	visitor := Visitor{}

	name, ok := vl["name"]
	if ok {
		visitor.Name = strings.Join(name, ",")
	}

	h.tmpl.Execute(w, visitor)
}

func NewHello(tmplPath string) (*Hello, error) {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return nil, err
	}

	return &Hello{tmpl}, nil
}

func main() {
	hello, err := NewHello("./index.html")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", hello)
	http.Handle(
		"/statics/",
		http.StripPrefix("/statics/", http.FileServer(http.Dir("./public"))),
	)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
