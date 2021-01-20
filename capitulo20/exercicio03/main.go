package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var contador int

func main() {
	incrementacao(60)
}

func incrementacao(x int) {
	wg.Add(x)
	for i := 0; i < x; i++ {
		go func() {
			y := contador
			runtime.Gosched()
			y++
			contador = y
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(x)
}
