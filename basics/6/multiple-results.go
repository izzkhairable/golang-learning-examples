package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func unswap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	unswappedA, unswappedB := swap(a, b)
	fmt.Println(a, b)
	fmt.Println(unswappedA, unswappedB)
}