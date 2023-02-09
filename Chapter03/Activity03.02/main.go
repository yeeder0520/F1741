package main

import (
	"errors"
	"fmt"
)

const (
	goodScore      = 450
	lowScoreRatio  = 10
	goodScoreRatio = 20
)

var (
	ErrCreditScore = errors.New("invalid credit score")
	ErrIncome      = errors.New("income invalid")
	ErrLoanAmount  = errors.New("loan amount invalid")
	ErrLoanTerm    = errors.New("loan term not a multiple of 12")
)

func checkLoan(creditScore int, income float64, loanAmount float64, loanTerm float64) error {
	// 信用良好者能得到較低的利率
	interest := 20.0
	if creditScore >= goodScore {
		interest = 15.0
	}
	// 檢查信用分數
	if creditScore < 1 {
		return ErrCreditScore
	}
	// 檢查收入
	if income < 1 {
		return ErrIncome
	}
	// 檢查貸款金額
	if loanAmount < 1 {
		return ErrLoanAmount
	}
	// 檢查貸款期數
	if loanTerm < 1 || int(loanTerm)%12 != 0 {
		return ErrLoanTerm
	}

	rate := interest / 100
	payment := ((loanAmount * rate) / loanTerm) + (loanAmount / loanTerm)

	// 計算貸款利息
	totalInterest := (payment * loanTerm) - loanAmount
	// 申請人能否負擔利息?
	approved := false
	if income > payment {
		ratio := (payment / income) * 100
		if creditScore >= goodScore && ratio < goodScoreRatio {
			approved = true
		} else if ratio < lowScoreRatio {
			approved = true
		}
	}

	fmt.Println("信用分數:", creditScore)
	fmt.Println("收入    :", income)
	fmt.Println("貸款金額:", loanAmount)
	fmt.Println("貸款利率:", interest)
	fmt.Println("貸款期數:", loanTerm)
	fmt.Println("總利息  :", totalInterest)
	fmt.Println("每月利息:", payment)
	fmt.Println("申請通過:", approved)
	fmt.Println("")

	return nil
}

func main() {
	// Approved
	fmt.Println("申請人 1")
	fmt.Println("-----------")
	err := checkLoan(500, 1000, 1000, 24)
	if err != nil {
		fmt.Println("錯誤:", err)
		return
	}
	// Denied
	fmt.Println("申請人 2")
	fmt.Println("-----------")
	err = checkLoan(350, 1000, 10000, 12)
	if err != nil {
		fmt.Println("錯誤:", err)
		return
	}
}
