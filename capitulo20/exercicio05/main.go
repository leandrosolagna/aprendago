package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var contador int64

func main() {
	incrementacao(60)
}

func incrementacao(x int) {
	wg.Add(x)
	for i := 0; i < x; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			v := atomic.LoadInt64(&contador)
			runtime.Gosched()
			fmt.Println(v)
			wg.Done()
		}()
	}

	wg.Wait()
}
