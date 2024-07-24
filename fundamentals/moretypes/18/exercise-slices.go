package main

import "golang.org/x/tour/pic"
import "fmt"

func Pic(dx, dy int) [][]uint8 {
	fmt.Println(dx, dy)
	a := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx, dx)
	}
	fmt.Println(a)
	// Do something.
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			switch {
			case j%2 == 0:
				a[i][j] = 150
			case i%8 == 0:
				a[i][j] = 5
			default:
				a[i][j] = 60
			}
		}
	}

	return a
}

func main() {
	pic.Show(Pic)
}
