package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

type Account struct {
	Savings, Current float64
}

var m map[string]Vertex
var a map[string]Account

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m)

	a = make(map[string]Account)
	a["Bob"] = Account{
		100,
		200,
	}
	fmt.Println(a["Bob"])
	fmt.Println(a)

}
