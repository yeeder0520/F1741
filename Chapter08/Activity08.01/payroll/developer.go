package payroll

import (
	"errors"
	"fmt"
)

var reviewStrToInt = map[string]int{
	"Excellent":      5,
	"Good":           4,
	"Fair":           3,
	"Poor":           2,
	"Unsatisfactory": 1,
}

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

func (d Developer) pay() (string, float64) {
	return d.Individual.fullName(), d.HourlyRate * d.HoursWorkedInYear
}

func (d Developer) ReviewRating() error {
	total := 0
	for _, v := range d.Review {
		rating, err := overallReview(v)
		if err != nil {
			return err
		}
		total += rating
	}
	averageRating := float64(total) / float64(len(d.Review))
	fmt.Printf("%s 今年考績評為 %.2f\n", d.Individual.fullName(), averageRating)
	return nil
}

func overallReview(i interface{}) (int, error) {
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
