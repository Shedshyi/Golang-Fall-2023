package Manager

import "Nurgalym/Employee"

type Manager struct {
	Employee.Employee
	position string
	salary   float32
	address  string
}

func (m *Manager) GetPosition() string {
	return m.position
}

func (m *Manager) GetSalary() float32 {
	return m.salary
}

func (m *Manager) GetAddress() string {
	return m.address
}

func (m *Manager) SetPosition(position string) {
	m.position = position
}

func (m *Manager) SetSalary(salary float32) {
	m.salary = salary
}

func (m *Manager) SetAddress(address string) {
	m.address = address
}
