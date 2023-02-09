package payroll

import "fmt"

type Payer interface {
	pay() (string, float64)
}

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

func (i Employee) fullName() string {
	return fmt.Sprintf("%v %v", i.FirstName, i.LastName)
}

func PayDetails(payers ...Payer) {
	for _, p := range payers {
		fullName, yearPay := p.pay()
		fmt.Printf("%s 今年薪資為 %.2f\n", fullName, yearPay)
	}
}
