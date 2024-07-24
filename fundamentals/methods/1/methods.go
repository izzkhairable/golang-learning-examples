package main

import (
	"fmt"
	"math"
	"strconv"
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

func (p Person) greet() string {
	greetMsg := "My name is "
	greetMsg += greetMsg
	greetMsg += p.name
	greetMsg += " my age is "
	greetMsg += strconv.Itoa(p.age)
	return greetMsg
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	p := Person{"Bob", 25}
	fmt.Println(p.greet())
}
