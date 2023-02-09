package main

import "fmt"

type salaryFunc func(int, int) int

func main() {
	devSalary := salary(50, 2080, developerSalary)
	bossSalary := salary(150000, 25000, managerSalary)

	fmt.Printf("經理薪資      : %d\n", bossSalary)
	fmt.Printf("程式設計師薪資: %d\n", devSalary)
}

func salary(x, y int, f salaryFunc) int {
	pay := f(x, y)
	return pay
}

func managerSalary(baseSalary, bonus int) int {
	return baseSalary + bonus
}

func developerSalary(hourlyRate, hoursWorked int) int {
	return hourlyRate * hoursWorked
}
