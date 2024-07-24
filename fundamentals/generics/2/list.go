package main

import (
	"fmt"
)

// List represents a singly-linked list that holds values of any type.
type List[T any] struct {
	val  T
	next *List[T]
}

// Insert adds a new value at the end of the list.
func (l *List[T]) Insert(val T) {
	if l.next == nil {
		l.next = &List[T]{val: val}
	} else {
		l.next.Insert(val)
	}
}

// Display prints all the values in the list.
func (l *List[T]) Display() {
	current := l
	for current != nil {
		fmt.Printf("%v ", current.val)
		current = current.next
	}
	fmt.Println()
}

func main() {
	// Create a list with an initial value
	list := &List[int]{val: 10}
	listTwo := &List[int]{val: 20}

	// Insert values into the list
	list.Insert(20)
	list.Insert(30)

	// Display the list
	fmt.Println("List after inserting 10, 20, 30:")
	list.Display()

	// Insert values into the list
	listTwo.Insert(100)
	listTwo.Insert(10)

	// Display the list
	fmt.Println("List after inserting 10, 20, 30:")
	listTwo.Display()
}
