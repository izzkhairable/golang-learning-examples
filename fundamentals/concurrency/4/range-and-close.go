package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
	// Error
	//c <- 100
}

func main() {
	c := make(chan int, 20)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
	d := c
	for i := range d {
		fmt.Println(i)
	}
}
