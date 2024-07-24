package main

import (
	"fmt"
	"math"
)

func main() {
	//var x, y int = 3, 4
	//var f float64 = math.Sqrt(float64(x*x + y*y))
	//var z uint = uint(f)
	//fmt.Println(x, y, z)
	var x, y = 3, 4
	xSquare := x * x
	fmt.Println(xSquare)
	ySquare := y * y
	fmt.Println(ySquare)
	var f = math.Sqrt(float64(xSquare + ySquare))
	var z = uint(f)
	fmt.Println(x, y, z)
}
