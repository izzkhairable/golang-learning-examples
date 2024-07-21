package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}

	var y [2]int
	fmt.Println(y, len(y), cap(y))
	if len(y) == 0 {
		fmt.Println("nil second!")
	} else {
		fmt.Println("ok")
	}
}
