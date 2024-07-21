package main

import "fmt"

type Vertex struct {
	X int
	Y int
	Z string
}

func main() {
	v := Vertex{1, 2, "Hello"}
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v.Y)
	fmt.Println(v.Z)
}
