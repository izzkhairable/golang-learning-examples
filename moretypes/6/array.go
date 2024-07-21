package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	fruits := [3]string{"apple", "watermelon", "orange"}
	fmt.Println(fruits)

	var pets [3]string
	pets = [3]string{"chicken", "dog", "cat"}
	fmt.Println(pets)
}
