package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"os"
)

func encrypt(data, key []byte) (resp []byte, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return resp, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return resp, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return resp, err
	}
	resp = gcm.Seal(nonce, nonce, data, nil)
	return resp, nil
}

func decrypt(data, key []byte) (resp []byte, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return resp, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return resp, err
	}
	nonce := data[:gcm.NonceSize()]
	encryptedData := data[gcm.NonceSize():]
	resp, err = gcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return resp, fmt.Errorf("解密錯誤: %v", err)
	}
	return resp, nil
}

func main() {
	data := "My secret text"
	fmt.Printf("原始資料: %s\n", data)

	key := make([]byte, 16)
	if _, err := rand.Read(key); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	encrypted, err := encrypt([]byte(data), key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("加密資料: %x\n", string(encrypted))

	decrypted, err := decrypt(encrypted, key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("解密資料: %s\n", string(decrypted))
}
