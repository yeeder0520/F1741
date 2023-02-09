package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
)

const (
	clientCertName = `.\client_cert.pem`
	serverCertName = `.\server_cert.pem`
	serverKeyName  = `.\server.key`
	host           = "localhost"
	port           = "8080"
)

func main() {
	clientCert, err := os.ReadFile(clientCertName)
	if err != nil {
		log.Fatal(err)
	}

	clientCAs, err := x509.SystemCertPool()
	if err != nil {
		clientCAs = x509.NewCertPool()
	}
	if ok := clientCAs.AppendCertsFromPEM(clientCert); !ok {
		log.Println("加入客戶端憑證錯誤")
	}

	tlsConfig := &tls.Config{
		ClientCAs:  clientCAs,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	server := &http.Server{
		Addr:      host + ":" + port,
		Handler:   nil,
		TLSConfig: tlsConfig,
	}

	log.Println("啟動伺服器")
	http.HandleFunc("/", hello)
	log.Fatal(server.ListenAndServeTLS(serverCertName, serverKeyName))
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("收到請求")
	w.Write([]byte("Hello Golang from a secure server"))
}
