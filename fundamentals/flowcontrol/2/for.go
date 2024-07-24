package main

import "fmt"

func main() {
	ascending()
	descending()
}

func ascending() {
	sum := 1
	for sum < 10 {
		fmt.Println(sum)
		sum += 1
	}
	fmt.Println(sum)
}

func descending() {
	sum := 10
	for sum > 0 {
		fmt.Println(sum)
		sum -= 1
	}
	fmt.Println(sum)
}
