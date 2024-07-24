package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func minus(x int, y int) int {
	return x - y
}
func multiply(x int, y int) int {
	return x * y
}

func stringAndInteger(name string) string {
	return "Hello" + name
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(minus(42, 13))
	fmt.Println(multiply(42, 13))
	fmt.Println(stringAndInteger("izzat"))
}
