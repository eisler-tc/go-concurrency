package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go doJob(&wg)
	}
	wg.Wait()
}

func doJob(wg *sync.WaitGroup) {

	time.Sleep(time.Second)
	fmt.Println("doing job")
	fmt.Println(wg)
	wg.Done()
}
