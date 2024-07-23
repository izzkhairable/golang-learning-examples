package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("error not able to squareroot the following: %v", float64(err))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(9))
	fmt.Println(Sqrt(-2))
	fmt.Println(Sqrt(-100))
}
