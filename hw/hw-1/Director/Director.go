package Director

import "Nurgalym/Employee"

type Director struct {
	Employee.Employee
	position string
	salary   float32
	address  string
}

func (d *Director) GetPosition() string {
	return d.position
}

func (d *Director) GetSalary() float32 {
	return d.salary
}

func (d *Director) GetAddress() string {
	return d.address
}

func (d *Director) SetPosition(position string) {
	d.position = position
}

func (d *Director) SetSalary(salary float32) {
	d.salary = salary
}

func (d *Director) SetAddress(address string) {
	d.address = address
}
