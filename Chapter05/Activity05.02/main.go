package main

import "fmt"

type Developer struct {
	Name          string
	HourlyRate    int
	WorkWeek      [7]int
	WorkNonLogged func(int) int
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
	d.WorkNonLogged = nonLoggedHours()

	fmt.Println("追蹤周一工時:", d.WorkNonLogged(8))
	logHours(&d, Monday)

	fmt.Println("追蹤周二工時:", d.WorkNonLogged(7))
	fmt.Println("追蹤周二工時:", d.WorkNonLogged(3))
	logHours(&d, Tuesday)

	fmt.Println("周一工時:", d.WorkWeek[Monday])
	fmt.Println("周二工時:", d.WorkWeek[Tuesday])
	fmt.Printf("本周薪資: %d", weekPayment(d))
}

func nonLoggedHours() func(int) int {
	total := 0
	return func(i int) int {
		if i < 0 {
			totalTmp := total
			total = 0
			return totalTmp
		}
		total += i
		return total
	}
}

func logHours(d *Developer, day Weekday) {
	d.WorkWeek[day] = d.WorkNonLogged(-1)
}

func weekPayment(d Developer) int {
	total := 0
	for _, v := range d.WorkWeek {
		total += v
	}
	return total * d.HourlyRate
}
