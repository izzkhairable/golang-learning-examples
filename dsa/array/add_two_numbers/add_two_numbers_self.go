package main

import (
	"strconv"
)

// AddTwoNumbersSelf solves the problem in O(n) time and O(1) space.
func AddTwoNumbersSelf(num1, num2 []int) []int {
	num1str := ""
	num2str := ""
	for _, v := range num1 {
		num1str += strconv.Itoa(v)
	}
	for _, v := range num2 {
		num2str += strconv.Itoa(v)
	}
	num1int, _ := strconv.Atoi(num1str)
	num2int, _ := strconv.Atoi(num2str)
	sum := num2int + num1int
	sumstr := strconv.Itoa(sum)
	var sumarr []int
	for _, ch := range sumstr {
		chInt, _ := strconv.Atoi(string(ch))
		sumarr = append(sumarr, chInt)
	}
	return sumarr

	//fmt.Println(num1str)
	//fmt.Println(num2str)
	//fmt.Println(sum)
	//fmt.Println(sumarr)
}
