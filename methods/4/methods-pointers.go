package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type Person struct {
	name string
	age  int
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (person *Person) update(name string, age int) {
	person.name = name
	person.age = age
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)

	v.Scale(20)
	fmt.Println(v)

	v.Abs()
	fmt.Println(v)

	person := Person{"Bob", 24}
	fmt.Println(person)

	person.update("Alice", 40)
	fmt.Println(person)
}
