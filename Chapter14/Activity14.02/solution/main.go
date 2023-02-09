package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Account struct {
	ID       string `json:id`
	Password string `json:password`
}

type LoginResult struct {
	Status string `json:status`
}

func loginServer(user Account, token string) LoginResult {
	jsonBytes, _ := json.Marshal(user)
	buffer := bytes.NewBuffer(jsonBytes)

	req, err := http.NewRequest("POST", "http://localhost:8080", buffer)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	result := LoginResult{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func main() {
	user := Account{ID: "John", Password: "1234"}
	token := "AuthorizedUser1234"
	result := loginServer(user, token)
	log.Printf("Login status: %v\n", result.Status)
}
