package main

import "fmt"

func main() {
	fmt.Println("my name")

	defer fmt.Println("world")

	fmt.Println("hello")
}
