package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const authorizedToken = "AuthorizedUser1234"

const (
	MsgServerInternalError = "server internal error"
	MsgInvalidToken        = "invalid authorization token"
	MsgInvalidID           = "invalid login ID"
	MsgInvalidPassword     = "invalid login password"
	MsgLogin               = "login successful"
)

var accounts = map[string]string{
	"John": "1234",
}

type Account struct {
	ID       string `json:id`
	Password string `json:password`
}

type LoginResult struct {
	Status string `json:status`
}

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logFile, err := os.OpenFile("login.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer logFile.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getJsonBytes(MsgServerInternalError))
		return
	}

	logger := log.New(logFile, "", log.Ldate|log.Lmicroseconds|log.Llongfile)
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	auth := r.Header.Get("Authorization")
	if auth != authorizedToken {
		log.Println(MsgInvalidToken)
		logger.Println(MsgInvalidToken)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(getJsonBytes(MsgInvalidToken))
		return
	}

	user := Account{}
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	pw, ok := accounts[user.ID]
	if !ok {
		log.Println(MsgInvalidID)
		logger.Println(MsgInvalidID)
		w.Write(getJsonBytes(MsgInvalidID))
		return
	} else if pw != user.Password {
		log.Println(MsgInvalidPassword)
		logger.Println(MsgInvalidPassword)
		w.Write(getJsonBytes(MsgInvalidPassword))
		return
	}

	successMsg := fmt.Sprintf(MsgLogin+": %v", user.ID)
	log.Println(successMsg)
	logger.Println(successMsg)
	w.Write(getJsonBytes(successMsg))
}

func getJsonBytes(msg string) []byte {
	jsonBytes, _ := json.Marshal(LoginResult{Status: msg})
	return jsonBytes
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
