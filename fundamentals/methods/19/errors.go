package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

type TestError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func (e *TestError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func runSecond() error {
	return &TestError{
		time.Now(),
		"it is not good",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
	if err := runSecond(); err != nil {
		fmt.Println(err)
	}
}
