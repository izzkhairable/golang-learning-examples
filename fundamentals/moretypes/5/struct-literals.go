package main

import "fmt"

type Vertex struct {
	X, Y int
}

type Employee struct {
	EmployeeName   string
	EmployeeID     string
	EmployeeSalary float64
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex

	z = &Vertex{212, 128}

	a = &Employee{"Bob", "E001", 1000.2}
	b = &Employee{EmployeeName: "Alice", EmployeeID: "M001"}
)

func main() {
	fmt.Println(v1, p, v2, v3, a, b)
}
