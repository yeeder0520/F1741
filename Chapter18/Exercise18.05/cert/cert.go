package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

const (
	clientCertName = `.\client_cert.pem`
	clientKeyName  = `.\client.key`
	serverCertName = `.\server_cert.pem`
	serverKeyName  = `.\server.key`
	host           = "127.0.0.1"
	hostDNS        = "localhost"
)

func main() {
	if err := generateCert(clientCertName, clientKeyName); err != nil {
		log.Println(err)
	}
	log.Println("產生:", clientCertName, clientKeyName)

	if err := generateCert(serverCertName, serverKeyName); err != nil {
		log.Println(err)
	}
	log.Println("產生:", serverCertName, serverKeyName)
}

func generateCert(certFile, keyFile string) error {
	serialNumber, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		return err
	}

	now := time.Now()

	ca := &x509.Certificate{
		Subject: pkix.Name{
			CommonName:    "Company",
			Organization:  []string{"Company, INC."},
			Country:       []string{"US"},
			Province:      []string{""},
			Locality:      []string{"San Francisco"},
			StreetAddress: []string{"Golden Gate Bridge"},
			PostalCode:    []string{"94016"},
		},
		SerialNumber:       serialNumber,
		SignatureAlgorithm: x509.SHA256WithRSA,
		// SignatureAlgorithm:    x509.ECDSAWithSHA256,
		NotBefore:             now,
		NotAfter:              now.AddDate(10, 0, 0),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IPAddresses:           []net.IP{net.ParseIP(host)},
		DNSNames:              []string{hostDNS},
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	// privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		return err
	}

	DER, err := x509.CreateCertificate(
		rand.Reader, ca, ca, &privateKey.PublicKey, privateKey)
	if err != nil {
		return err
	}

	cert := pem.EncodeToMemory(
		&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: DER,
		})

	/*
		pemByte, err := x509.MarshalECPrivateKey(privateKey)
		if err != nil {
			return err
		}
	*/
	key := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
			// Type: "PRIVATE KEY",
			// Bytes: pemByte,
		})

	if err := os.WriteFile(certFile, cert, 0777); err != nil {
		return err
	}
	if err := os.WriteFile(keyFile, key, 0600); err != nil {
		return err
	}

	return nil
}
