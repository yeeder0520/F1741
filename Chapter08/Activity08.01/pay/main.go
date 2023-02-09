package main

import (
	"fmt"
	"os"

	pr "Activity08.01/payroll"
)

var (
	employeeReview = make(map[string]interface{})
	d              pr.Developer
	m              pr.Manager
)

func init() {
	fmt.Println("初始化變數...")
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"

	d = pr.Developer{
		Individual:        pr.Employee{Id: 1, FirstName: "Eric", LastName: "Davis"},
		HourlyRate:        35,
		HoursWorkedInYear: 2400,
		Review:            employeeReview,
	}
	m = pr.Manager{
		Individual:     pr.Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"},
		Salary:         150000,
		CommissionRate: .07,
	}
}

func main() {
	err := d.ReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pr.PayDetails(d, m)
}
