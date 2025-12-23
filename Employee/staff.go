package Employee

import "fmt"

type Staff struct {
	Employees []Payable
}

// Add employee to staff
func (s *Staff) AddEmployee(emp Payable) {
	s.Employees = append(s.Employees, emp)
}

// Print salaries
func (s Staff) PrintSalaries() {
	for _, emp := range s.Employees {
		fmt.Printf(
			"Employee: %s | Salary: %.2f\n",
			emp.GetName(),
			emp.CalculateSalary(),
		)
	}
}
