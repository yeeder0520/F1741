package main

import "fmt"

type Developer struct {
	Name       string
	HourlyRate int
	WorkWeek   [7]int
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	d := Developer{Name: "Tony Stark", HourlyRate: 10}
	d.logHours(Monday, 8)
	d.logHours(Tuesday, 10)
	fmt.Println("周一工時:", d.WorkWeek[Monday])
	fmt.Println("周二工時:", d.WorkWeek[Tuesday])
	fmt.Printf("本周薪資: %d", d.weekPayment())
}

func (d *Developer) logHours(day Weekday, hours int) {
	d.WorkWeek[day] = hours
}

func (d Developer) weekPayment() int {
	total := 0
	for _, v := range d.WorkWeek {
		total += v
	}
	return total * d.HourlyRate
}
