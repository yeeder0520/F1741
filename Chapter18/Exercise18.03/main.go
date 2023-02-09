package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	data := []byte("My secret text")
	fmt.Printf("原始資料: %s\n", data)

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("產生私鑰錯誤: %v", err)
		os.Exit(1)
	}
	publicKey := privateKey.PublicKey

	encrypted, err := rsa.EncryptOAEP(
		sha256.New(), rand.Reader, &publicKey, data, nil)
	if err != nil {
		fmt.Printf("加密錯誤: %v", err)
		os.Exit(1)
	}
	fmt.Printf("加密資料: %x\n", string(encrypted))

	decrypted, err := rsa.DecryptOAEP(
		sha256.New(), rand.Reader, privateKey, encrypted, nil)
	if err != nil {
		fmt.Printf("解密錯誤: %v", err)
		os.Exit(1)
	}
	fmt.Printf("解密資料: %s\n", string(decrypted))
}
