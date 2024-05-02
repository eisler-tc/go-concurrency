package main

import (
	"fmt"
)

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	fmt.Println("Closing channel")
	defer close(ch)
	fmt.Println("Channel closed..")
}

func main() {
	ch := make(chan int)

	go producer(ch)
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Recieved ", v, ok)
	}
}
