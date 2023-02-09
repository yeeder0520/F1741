package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"os"
)

func main() {
	data := []byte("My secret document")

	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	signedData := ed25519.Sign(privateKey, data)
	fmt.Printf("數位簽章:\n%x\n", signedData)

	verified := ed25519.Verify(publicKey, data, signedData)
	fmt.Println("驗證:", verified)
}
