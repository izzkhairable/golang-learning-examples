package main

import "fmt"

func main() {
	ascending()
	descending()
}

func ascending() {
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)
}

func descending() {
	sum := 0
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)
}
