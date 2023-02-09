package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	ErrInvalidSSNLength     = errors.New("SSN 不足 9 字元")
	ErrInvalidSSNNumbers    = errors.New("SSN 非數字")
	ErrInvalidSSNPrefix     = errors.New("SSN 以 000 開頭")
	ErrInvalidSSNDigitPlace = errors.New("SSN 以 9 開頭時第 4 位需為 7 或 9")
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	validateSSN := []string{"123-45-6789", "012-8-678", "000-12-0962", "999-33-3333", "087-65-4321", "123-45-zzzz"}

	for idx, ssn := range validateSSN {
		log.Printf("驗證資料 %d 之 %d: %#v ", len(validateSSN), idx+1, ssn)
		ssn = strings.Replace(ssn, "-", "", -1)
		err := isNumber(ssn)
		if err != nil {
			log.Println(err)
		}
		err = validLength(ssn)
		if err != nil {
			log.Println(err)
		}
		err = isPrefixValid(ssn)
		if err != nil {
			log.Println(err)
		}
		err = validDigitPlace(ssn)
		if err != nil {
			log.Println(err)
		}
	}
}

func validLength(ssn string) error {
	ssn = strings.TrimSpace(ssn)
	if len(ssn) != 9 {
		return fmt.Errorf("值 %s 導致錯誤: %v", ssn, ErrInvalidSSNLength)
	}
	return nil
}

func isNumber(ssn string) error {
	_, err := strconv.Atoi(ssn)
	if err != nil {
		return fmt.Errorf("值 %s 導致錯誤: %v", ssn, ErrInvalidSSNNumbers)
	}
	return nil
}

func isPrefixValid(ssn string) error {
	if strings.HasPrefix(ssn, "000") {
		return fmt.Errorf("值 %s 導致錯誤: %v", ssn, ErrInvalidSSNPrefix)
	}
	return nil
}

func validDigitPlace(ssn string) error {
	if string(ssn[0]) == "9" && (string(ssn[3]) != "9" && string(ssn[3]) != "7") {
		return fmt.Errorf("值 %s 導致錯誤: %v", ssn, ErrInvalidSSNDigitPlace)
	}
	return nil
}
