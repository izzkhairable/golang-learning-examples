package main

import (
	"fmt"
	"math"
)

type MyFloat float64
type AnotherFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f AnotherFloat) Ceil() float64 {
	return math.Ceil(float64(f))
}

func (f AnotherFloat) GetFloat() float64 {
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	z := AnotherFloat(9.235)
	fmt.Println(z)
	fmt.Println(z.Ceil())
	fmt.Println(z.GetFloat())
}
