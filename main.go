// package main

// import "fmt"

// type Employee struct {
// 	Name   string
// 	Salary float64
// }

// func (e *Employee) increaseSalary(percentage float64) {
// 	e.Salary += e.Salary * percentage / 100
// }

// func main() {
// 	emp := Employee{Name: "John", Salary: 5000}
// 	emp.increaseSalary(10)
// 	fmt.Println(emp.Salary)

// }

// Soal no 4

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func printNumbers() {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println(i)
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }

// func main() {
// 	go printNumbers()
// 	fmt.Println("selesai")
// }

// soal 5

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "go is amaring"
	words := strings.Split(s, " ")
	fmt.Println(len(words))
}
