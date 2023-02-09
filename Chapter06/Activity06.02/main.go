package main

import (
	"errors"
	"fmt"
	"strings"
)

type directDeposit struct {
	lastName      string
	firstName     string
	bankName      string
	routingNumber int
	accountNumber int
}

var ErrInvalidLastName = errors.New("invalid last name")
var ErrInvalidRoutingNum = errors.New("invalid routing number")

func main() {
	dd := directDeposit{
		lastName:      "  ",
		firstName:     "Abe",
		bankName:      "XYZ Inc",
		routingNumber: 17,
		accountNumber: 1809,
	}

	err := dd.validateRoutingNumber()
	if err != nil {
		fmt.Println(err)
	}
	err = dd.validateLastName()
	if err != nil {
		fmt.Println(err)
	}

	dd.report()

}

func (dd *directDeposit) validateRoutingNumber() error {
	if dd.routingNumber < 100 {
		return ErrInvalidRoutingNum
	}
	return nil
}

func (dd *directDeposit) validateLastName() error {
	dd.lastName = strings.TrimSpace(dd.lastName)
	if len(dd.lastName) == 0 {
		return ErrInvalidLastName
	}
	return nil
}

func (dd *directDeposit) report() {
	fmt.Println("--------------------")
	fmt.Println("姓:", dd.lastName)
	fmt.Println("名:", dd.firstName)
	fmt.Println("銀行名稱:", dd.bankName)
	fmt.Println("轉帳代碼: ", dd.routingNumber)
	fmt.Println("帳號: ", dd.accountNumber)
}
