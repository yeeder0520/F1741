package main

import (
	"log"
	"net/http"
)

const (
	serverCertName = `.\server_cert.pem`
	serverKeyName  = `.\server.key`
)

func main() {
	log.Println("啟動伺服器")
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServeTLS(":8080", serverCertName, serverKeyName, nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("收到請求")
	w.Write([]byte("Hello Golang from a secure server"))
}
