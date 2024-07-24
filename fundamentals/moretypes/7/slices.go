package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	var z []int = primes[3:5]
	var first = primes[0]
	var last = primes[5]
	fmt.Println(s, z, first, last)
}
