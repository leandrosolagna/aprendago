package main

import "fmt"

func main() {
	canal := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				canal <- j
			}
		}()
	}
	demonstrar(canal)
}

func demonstrar(c chan int) {
	for k := 0; k < 100; k++ {
		fmt.Println(<-c)
	}
}
