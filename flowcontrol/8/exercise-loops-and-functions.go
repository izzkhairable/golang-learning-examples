package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for z < 10 {
		fmt.Println(z)
		z = (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
