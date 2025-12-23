package Employee

// Interface
type Payable interface {
	CalculateSalary() float64
	GetName() string
}

// Base struct
type Employee struct {
	ID         int
	Name       string
	Position   string
	BaseSalary float64
}

// Method
func (e Employee) CalculateSalary() float64 {
	return e.BaseSalary
}

func (e Employee) GetName() string {
	return e.Name
}
