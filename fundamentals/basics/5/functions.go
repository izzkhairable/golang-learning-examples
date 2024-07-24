package main

import (
	"fmt"
	"strconv"
)

func add(x, y int) int {
	return x + y
}

func stringAndInteger(name string, age, salary int) string {
	return "Hello " + name + " " + strconv.Itoa(age) + strconv.Itoa(salary)
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(stringAndInteger("izzat", 70, 1000))
}
