package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
		//		q <- 0
	}()
	return c
}

func receive(c <-chan int, q chan int) {
	for {
		select {
		case v, ok := <-c:
			if ok {
				fmt.Println(v)
			} else {
				return
			}
		case v, ok := <-q:
			if ok {
				fmt.Println(v)
			} else {
				return
			}
		}
	}
}
