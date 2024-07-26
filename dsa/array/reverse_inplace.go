package main

import "fmt"

// ReverseInPlace solves the problem in O(n) time and O(1) space.
func ReverseInPlace(list []int, start, end int) {
	fmt.Printf("\nlist []int = %d", list)
	fmt.Printf("\nstart = %d", start)
	fmt.Printf("\nend = %d", end)
	fmt.Println("\n--------------------------")
	fmt.Println()
	fmt.Printf("  i := start; i <= start+end/2; i++\n")
	fmt.Printf("  i := %d;\n  i <= %d + %d / 2\n  && i < %d -i +%d; i++\n", start, start, end, end, start)

	//for i := start; i <= start+end/2 && i < end-i+start; i++
	for i := start; i <= start+end/2 && i < end-i+start; i++ {

		fmt.Printf("\n   i := %d \n", i)
		fmt.Printf("\n   start+end/2 := %d \n", start+end/2)
		fmt.Printf("\n   end-i+start := %d \n", end-i+start)
		fmt.Printf("   list[i], list[end-i+start] = list[end-i+start], list[i] \n")
		fmt.Printf("   list[%d], list[%d-%d+%d] = list[%d-%d+%d], list[%d]", i, end, i, start, end, i, start, i)

		list[i], list[end-i+start] = list[end-i+start], list[i]
		fmt.Println("")
	}
	fmt.Println()
	fmt.Println("--------------------------")
	fmt.Println(list)
}

func ReverseInPlacePractice(list []int, start int, end int) {
	for i := start; i <= start+end/2; i++ {
		//frontInt := list[i]
		//backInt := list[i-2]
	}
}
