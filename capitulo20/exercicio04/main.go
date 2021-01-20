package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var contador int
var mtx sync.Mutex

func main() {
	incrementacao(60)
}

func incrementacao(x int) {
	wg.Add(x)
	for i := 0; i < x; i++ {
		go func() {
			mtx.Lock()
			y := contador
			runtime.Gosched()
			y++
			contador = y
			mtx.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(contador)
}
