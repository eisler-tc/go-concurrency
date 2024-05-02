package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//create a wait group
	wg := sync.WaitGroup{}

	//create a channel which can hold 1 element
	sharedNumber := make(chan int, 10)

	//listen for data coming from sharedNumber channel
	wg.Add(10)
	go func() {
		for i := 0; i < 10; i++ {
			//receive data from shared number channel
			time.Sleep(time.Second)
			received := <-sharedNumber
			fmt.Printf("I received: %v.\n", received)
			wg.Done()
		}

	}()

	wg.Add(1)
	//Create and send data to sharedNumber channel
	go func() {
		for i := 0; i < 10; i++ {
			//Send i to sharedNumber channel
			sharedNumber <- i
			fmt.Println("I sent the data..")
		}
		wg.Done()
	}()

	//Wait for all go routines are done
	wg.Wait()
}
