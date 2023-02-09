package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("無效的時薪")
	ErrHoursWorked = errors.New("無效的一周工時")
)

func main() {
	pay := payDay(100, 25)
	fmt.Printf("周薪: %d\n\n", pay)

	pay = payDay(100, 200)
	fmt.Printf("周薪: %d\n\n", pay)

	pay = payDay(60, 25)
	fmt.Printf("周薪: %d\n\n", pay)
}

func payDay(hoursWorked, hourlyRate int) int {
	defer func() {
		if r := recover(); r != nil {
			if r == ErrHourlyRate {
				fmt.Printf("時薪: %d\n錯誤: %v\n", hourlyRate, r)
			}
			if r == ErrHoursWorked {
				fmt.Printf("工時: %d\n錯誤: %v\n", hoursWorked, r)
			}
		}
		fmt.Printf("計算周薪的依據: 工時: %d / 時薪: %d\n", hoursWorked, hourlyRate)
	}()
	if hourlyRate < 10 || hourlyRate > 75 {
		panic(ErrHourlyRate)
	}

	if hoursWorked < 0 || hoursWorked > 80 {
		panic(ErrHoursWorked)
	}

	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overTime
	}
	return hoursWorked * hourlyRate
}
