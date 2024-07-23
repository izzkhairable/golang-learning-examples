package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
	//
	//f = i.(float64) // panic
	//fmt.Println(f)

	var myInterface interface{} = 2.86
	val := myInterface.(float64)
	fmt.Println(val)

	valtwo, ok := myInterface.(string)
	fmt.Println(valtwo, ok)
}
