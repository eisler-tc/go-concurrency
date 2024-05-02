package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var a int32

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1000000)
	for i := 0; i < 1000000; i++ {
		go func() {
			atomic.AddInt32(&a, 1)
			wg.Done()
		}()
	}
	wg.Add(1000000)
	for i := 0; i < 1000000; i++ {
		go func() {
			atomic.AddInt32(&a, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(a)
}
