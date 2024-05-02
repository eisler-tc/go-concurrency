package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type SharedInteger struct {
	mu    sync.Mutex
	value int32
}

func (sharedInteger *SharedInteger) Increment() {
	atomic.AddInt32(&sharedInteger.value, 1)
}

func main() {

	wg := sync.WaitGroup{}
	wg.Add(10000)

	sharedInteger := SharedInteger{}
	wg.Add(-9999)
	for i := 0; i < 10000; i++ {
		go func() {
			sharedInteger.Increment()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("SharedInteger=%v", sharedInteger.value)
}
