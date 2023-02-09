package payroll

type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

func (m Manager) pay() (string, float64) {
	return m.Individual.fullName(), m.Salary + (m.Salary * m.CommissionRate)
}
