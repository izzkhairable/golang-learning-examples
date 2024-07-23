package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	myImplementation()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type myDb interface {
	getRecords() string
	updateRecords() string
	deleteRecords() string
	createRecords() string
}

type OrderDb struct {
	Row    int
	Column int
}

func (orderDb *OrderDb) getRecords() string {
	return "Hello"
}

func (orderDb *OrderDb) updateRecords() string {
	return "Hello"
}

func (orderDb *OrderDb) deleteRecords() string {
	return "Hello"
}

func (orderDb *OrderDb) createRecords() string {
	return "Hello"
}

func myImplementation() {
	var myOrderDB myDb
	myOrderDB = &OrderDb{10, 10}
	fmt.Println(myOrderDB)
}
