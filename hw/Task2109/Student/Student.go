package Student

import (
	"Task/User"
)

type Student struct {
	User.User
	Id       string
	Name     string
	Surname  string
	userType string
}

func (s *Student) GetId() string {
	return s.Id
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) GetSurname() string {
	return s.Surname
}

func (s *Student) GetUserType() string {
	return "Student"
}
