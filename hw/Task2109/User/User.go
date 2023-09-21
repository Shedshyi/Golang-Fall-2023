package User

type User interface {
	GetId() string
	GetUserType() string
	GetName() string
	GetSurname() string
}
