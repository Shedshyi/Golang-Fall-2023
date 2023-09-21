package Worker

import (
	"Task/User"
)

type Worker struct {
	User.User
	Id       string
	Name     string
	Surname  string
	userType string
}

func (w *Worker) GetId() string {
	return w.Id
}

func (w *Worker) GetName() string {
	return w.Name
}

func (w *Worker) GetSurname() string {
	return w.Surname
}

func (w *Worker) GetUserType() string {
	return "Worker"
}
