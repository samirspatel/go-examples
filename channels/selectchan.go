package main

import "fmt"

func fibonacci2(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y := y, x+y
			fmt.Printf("x: %d, y: %d \n", x, y)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func ChanSelect() {
	fmt.Println("ChanSelect")
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}
