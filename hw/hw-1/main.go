package main

import (
	"Nurgalym/Accountant"
	"Nurgalym/Director"
	"Nurgalym/Engineer"
	"Nurgalym/Hr"
	"Nurgalym/Manager"
	"fmt"
)

func main() {
	m := Manager.Manager{}
	e := Engineer.Engineer{}
	d := Director.Director{}
	h := Hr.Hr{}
	a := Accountant.Accountant{}

	a.SetPosition("Seo Director")
	e.SetPosition("Tech Safety Engineer")
	m.SetPosition("Sales Manager")
	h.SetPosition("Hr Seo")
	d.SetPosition("Tech Director")

	m.SetSalary(550000.050)
	a.SetSalary(450000.050)
	d.SetSalary(850290.050)
	h.SetSalary(650000.050)
	e.SetSalary(550000.050)

	m.SetAddress("Tole bi 59")
	d.SetAddress("Tole bi 61")
	e.SetAddress("Tole bi 63")
	h.SetAddress("Tole bi 65")
	a.SetAddress("Tole bi 67")

	fmt.Println("Manager:\nSalary: ", m.GetSalary(), "\nPosition:", m.GetPosition(), "\nAddress:", m.GetAddress(), "\n")
	fmt.Println("Engineer:\nSalary: ", e.GetSalary(), "\nPosition:", e.GetPosition(), "\nAddress:", e.GetAddress(), "\n")
	fmt.Println("Hr:\nSalary: ", h.GetSalary(), "\nPosition:", h.GetPosition(), "\nAddress:", h.GetAddress(), "\n")
	fmt.Println("Director\nSalary: ", d.GetSalary(), "\nPosition:", d.GetPosition(), "\nAddress:", d.GetAddress(), "\n")
	fmt.Println("Accountant:\nSalary: ", a.GetSalary(), "\nPosition:", a.GetPosition(), "\nAddress:", a.GetAddress(), "\n")

}
