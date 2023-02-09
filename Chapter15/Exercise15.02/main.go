package main

import (
	"log"
	"net/http"
)

type hello struct{}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hello World</h1>"
	w.Write([]byte(msg))
}

func servePage1(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Page 1</h1>"
	w.Write([]byte(msg))
}

func main() {
	http.HandleFunc("/page1", servePage1)
	http.Handle("/", hello{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
