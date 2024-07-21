package main

import "fmt"

func main() {
	i, j, z := 42, 2701, 300

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	y := &z
	*y = *y * i * j
	fmt.Println(z)

}
