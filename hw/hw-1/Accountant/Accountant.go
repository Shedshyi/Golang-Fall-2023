package Accountant

import "Nurgalym/Employee"

type Accountant struct {
	Employee.Employee
	position string
	salary   float32
	address  string
}

func (a *Accountant) GetPosition() string {
	return a.position
}

func (a *Accountant) GetSalary() float32 {
	return a.salary
}

func (a *Accountant) GetAddress() string {
	return a.address
}

func (a *Accountant) SetPosition(position string) {
	a.position = position
}

func (a *Accountant) SetSalary(salary float32) {
	a.salary = salary
}

func (a *Accountant) SetAddress(address string) {
	a.address = address
}
