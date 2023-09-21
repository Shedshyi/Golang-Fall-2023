package main

import (
	"Task/Student"
	"Task/Worker"
	"fmt"
)

func main() {
	s := Student.Student{}
	fmt.Scan(&s.Surname)
	fmt.Println(s.Surname)
	fmt.Println(s.GetUserType())

	w := Worker.Worker{}
	fmt.Scan(&w.Name, &w.Surname)
	fmt.Println(w.GetUserType())
}
