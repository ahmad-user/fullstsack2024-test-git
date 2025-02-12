package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func (e *Employee) increaseSalary(percentage float64) {
	e.Salary += e.Salary + percentage/100
}

func main() {
	emp := Employee{Name: "John", Salary: 5000}
	emp.increaseSalary(10)
	fmt.Println(emp.Salary)

}
