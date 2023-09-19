package Employee

type Employee interface {
	GetPosition() string
	SetPosition(position string)
	GetSalary() float32
	SetSalary(salary float32)
	GetAddress() string
	SetAddress(address string)
}
