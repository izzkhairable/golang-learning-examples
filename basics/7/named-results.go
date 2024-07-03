package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func splitThree(sum int) (sumReturned, x, y, z int) {
	sumReturned = sum
	x = sum * 4 / 9
	y = sum - x
	z = sum - y
	return
}

func main() {
	fmt.Println(split(17))
	fmt.Println(splitThree(17))

}
