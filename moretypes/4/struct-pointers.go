package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	p.Y = 200

	m := &v
	m.X = 10
	m.Y = 100

	fmt.Println(v)
}
