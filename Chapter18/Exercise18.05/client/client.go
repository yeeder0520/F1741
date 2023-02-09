package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	clientCertName = `.\client_cert.pem`
	clientKeyName  = `.\client.key`
	serverCertName = `.\server_cert.pem`
	host           = "localhost"
	port           = "8080"
)

func main() {
	cert, err := tls.LoadX509KeyPair(clientCertName, clientKeyName)
	if err != nil {
		log.Fatal(err)
	}

	serverCert, err := os.ReadFile(serverCertName)
	if err != nil {
		log.Fatal(err)
	}

	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		rootCAs = x509.NewCertPool()
	}
	if ok := rootCAs.AppendCertsFromPEM(serverCert); !ok {
		log.Fatal("加入伺服器憑證錯誤")
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      rootCAs,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	resp, err := client.Get("https://" + host + ":" + port)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("收到回應:", string(data))
}
