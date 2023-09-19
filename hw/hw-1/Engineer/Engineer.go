package Engineer

import "Nurgalym/Employee"

type Engineer struct {
	Employee.Employee
	position string
	salary   float32
	address  string
}

func (e *Engineer) GetPosition() string {
	return e.position
}

func (e *Engineer) GetSalary() float32 {
	return e.salary
}

func (e *Engineer) GetAddress() string {
	return e.address
}

func (e *Engineer) SetPosition(position string) {
	e.position = position
}

func (e *Engineer) SetSalary(salary float32) {
	e.salary = salary
}

func (e *Engineer) SetAddress(address string) {
	e.address = address
}
