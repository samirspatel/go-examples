package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to chan c
}

// https://go.dev/tour/concurrency/4
func RunChan() {
	fmt.Println("RunChan")
	c := make(chan int)
	m := []int{7, 2, 8, -9, 4, 0, 123, 23, 123, 23, 423, 4234, 234}
	go sum(m[len(m)/2:], c)

	s := []int{7, 2, 8, -9, 4, 0}
	go sum(s[:len(s)/2], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

// ----

func fibonacci(n int, c chan int){
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// Note: Only the sender should close a channel, never the receiver. 
	// Sending on a closed channel will cause a panic.
	// Another note: Channels aren't like files; you don't usually need to close them. 
	// Closing is only necessary when the receiver must be told there are no more values coming, 
	// such as to terminate a range loop.
	close(c)
}

func RangeAndCloseChan(){
	fmt.Println("RangeAndCloseChan")
	c := make(chan int, 20)
	go fibonacci(cap(c), c)
	// The loop for i := range c receives values from the channel repeatedly until it is closed.
	for i := range c {
		fmt.Printf("%d ", i)
	}
}