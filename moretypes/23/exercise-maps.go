package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	fmt.Println(s)
	m := make(map[string]int)
	arrOfStr := strings.Split(s, " ")
	for _, v := range arrOfStr {
		m[v] += 1
	}
	//return map[string]int{"x": 1}
	fmt.Println(m)
	return m
}

func main() {
	wc.Test(WordCount)
}
