package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	incrementacao(60)
}

func incrementacao(x int) {
	wg.Add(x)
	for i := 0; i < x; i++ {
		go func() {
			y := 0
			runtime.Gosched()
			y++
			y = y
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(x)
}
