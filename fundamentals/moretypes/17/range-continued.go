package main

import "fmt"

func main() {
	pow := make([]int, 10)
	fmt.Println(pow)
	fmt.Println("-----")
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	fmt.Println(pow)
	fmt.Println("-----")
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	fmt.Println("-----")
	for index, _ := range pow {
		fmt.Printf("%d\n", index)
	}
	fmt.Println("-----")
}
