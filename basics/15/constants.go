package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	const Name = "Bob the builder"
	const Salary = 20001.62
	// cannot do this
	//const Salary := 20001.62
	fmt.Println("Name is ", Name, " Salary ", Salary)

	pet := "dog"
	pet = "cat"
	fmt.Println(pet)
}
