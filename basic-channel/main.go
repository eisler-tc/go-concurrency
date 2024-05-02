package main

import (
	"fmt"
	"sync"
)

func main() {

	//create a wait group
	wg := sync.WaitGroup{}

	//create a channel which can hold 1 element
	sharedNumber := make(chan int)

	//listen for data coming from sharedNumber channel
	go func() {
		for i := 0; i < 10; i++ {
			//receive data from shared number channel
			received := <-sharedNumber
			fmt.Printf("I received: %v.\n", received)
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
