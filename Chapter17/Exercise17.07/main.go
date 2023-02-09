package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux" // go get github.com/gorilla/mux
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<h1>Hello Golang with gorilla/mux</h1>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", exampleHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
