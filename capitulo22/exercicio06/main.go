package main

import "fmt"

func main() {
	canal := make(chan int)
	go loopNumeros(canal)
	demonstrar(canal)
}

func loopNumeros(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func demonstrar(r <-chan int) {
	for v := range r {
		fmt.Println(v)
	}
}
