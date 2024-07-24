package array

import "fmt"

// ReverseInPlace solves the problem in O(n) time and O(1) space.
func ReverseInPlace(list []int, start, end int) {
	for i := start; i <= start+end/2 && i < end-i+start; i++ {
		fmt.Println(start + end/2)
		fmt.Println(end - i + start)
		list[i], list[end-i+start] = list[end-i+start], list[i]
	}
}
