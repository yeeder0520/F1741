package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	c, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	time.Sleep(time.Second * 2)

	select {
	case <-c.Done():
		fmt.Println("Server timeout")
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("<h1>Server timeout</h1>"))
		return
	default:
		fmt.Println("Hello Golang")
		w.Write([]byte("<h1>Hello Golang</h1>"))
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
