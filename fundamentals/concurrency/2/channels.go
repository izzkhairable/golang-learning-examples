package main

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func msgHandle(msg string, broker chan string) {
	broker <- msg
}

func main() {
	//s := []int{7, 2, 8, -9, 4, 0}
	//
	//c := make(chan int)
	//go sum(s[:len(s)/2], c)
	//go sum(s[len(s)/2:], c)
	//x, y := <-c, <-c // receive from c
	//
	//fmt.Println(x, y, x+y)

	broker := make(chan string)
	go msgHandle("world", broker)
	go msgHandle("hello", broker)
	first, second := <-broker, <-broker

	print(first, " ", second)
}
