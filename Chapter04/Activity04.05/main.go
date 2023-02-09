package main

import (
	"fmt"
	"os"
	"strings"
)

type locale struct {
	language  string
	territory string
}

func getLocales() map[locale]bool {
	supportedLocales := make(map[locale]bool, 5)
	supportedLocales[locale{"en", "US"}] = true
	supportedLocales[locale{"en", "CN"}] = true
	supportedLocales[locale{"fr", "CN"}] = true
	supportedLocales[locale{"fr", "FR"}] = true
	supportedLocales[locale{"ru", "RU"}] = true
	return supportedLocales
}

func localeExists(l locale) bool {
	_, exists := getLocales()[l]
	return exists
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("沒有傳入語系")
		os.Exit(1)
	}

	localeParts := strings.Split(os.Args[1], "_")
	if len(localeParts) != 2 {
		fmt.Printf("語系輸入錯誤: %v\n", os.Args[1])
		os.Exit(1)
	}

	passedLocale := locale{
		territory: localeParts[1],
		language:  localeParts[0],
	}

	if !localeExists(passedLocale) {
		fmt.Printf("不支援傳入的語系: %v\n", os.Args[1])
		os.Exit(1)
	}
	fmt.Println("支援傳入的語系")
}
