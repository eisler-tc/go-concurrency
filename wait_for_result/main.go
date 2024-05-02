package main

import (
	"fmt"
	"time"
)

func main() {

	stopCh := make(chan interface{})

	go func() {
		time.Sleep(3 * time.Second)
		stopCh <- struct{}{}
	}()

	fmt.Println("I am doing some job.. Waiting for any signal from stop channel..")
	<-stopCh
	fmt.Println("Finally stop signal. Bye..")

}
