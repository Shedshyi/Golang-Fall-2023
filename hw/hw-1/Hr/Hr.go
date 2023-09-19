package Hr

import "Nurgalym/Employee"

type Hr struct {
	Employee.Employee
	position string
	salary   float32
	address  string
}

func (h *Hr) GetPosition() string {
	return h.position
}

func (h *Hr) GetSalary() float32 {
	return h.salary
}

func (h *Hr) GetAddress() string {
	return h.address
}

func (h *Hr) SetPosition(position string) {
	h.position = position
}

func (h *Hr) SetSalary(salary float32) {
	h.salary = salary
}

func (h *Hr) SetAddress(address string) {
	h.address = address
}
