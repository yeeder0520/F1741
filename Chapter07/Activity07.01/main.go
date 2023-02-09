package main

import (
	"errors"
	"fmt"
	"os"
)

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

type Payer interface {
	Pay() (string, float64)
}

var reviewStrToInt = map[string]int{
	"Excellent":      5,
	"Good":           4,
	"Fair":           3,
	"Poor":           2,
	"Unsatisfactory": 1,
}

func main() {
	employeeReview := make(map[string]interface{})
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"

	d := Developer{
		Individual:        Employee{Id: 1, FirstName: "Eric", LastName: "Davis"},
		HourlyRate:        35,
		HoursWorkedInYear: 2400,
		Review:            employeeReview,
	}
	m := Manager{
		Individual:     Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"},
		Salary:         150000,
		CommissionRate: .07,
	}

	err := d.ReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	payDetails(d, m)
}

func payDetails(payers ...Payer) {
	for _, p := range payers {
		fullName, yearPay := p.Pay()
		fmt.Printf("%s 今年薪資為 %.2f\n", fullName, yearPay)
	}
}

func (i Employee) FullName() string {
	return fmt.Sprintf("%v %v", i.FirstName, i.LastName)
}

func (m Manager) Pay() (string, float64) {
	return m.Individual.FullName(), m.Salary + (m.Salary * m.CommissionRate)
}

func (d Developer) Pay() (string, float64) {
	return d.Individual.FullName(), d.HourlyRate * d.HoursWorkedInYear
}

func (d Developer) ReviewRating() error {
	total := 0
	for _, v := range d.Review {
		rating, err := OverallReview(v)
		if err != nil {
			return err
		}
		total += rating
	}
	averageRating := float64(total) / float64(len(d.Review))
	fmt.Printf("%s 今年考績評為 %.2f\n", d.Individual.FullName(), averageRating)
	return nil
}

func OverallReview(i interface{}) (int, error) {
	switch v := i.(type) {
	case int:
		return v, nil
	case string:
		if rating, ok := reviewStrToInt[v]; ok {
			return rating, nil
		}
		return 0, errors.New("無效的績效評分: " + v)
	default:
		return 0, errors.New("無效的績效資料")
	}
}
