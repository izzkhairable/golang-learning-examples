package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
		return lim
	}
	// can't use v here, though

}

func validateStrLen(str string) string {
	if len(str) > 10 {
		return "ok good to go"
	} else {
		return "nah uh do better"
	}

}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
		validateStrLen("Hello how are you"),
	)
}
