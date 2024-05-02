package main

import (
	"fmt"
	"runtime"
)

func FanOut(message string) chan string {
	outChannel := make(chan string)
	go func() {
		outChannel <- message
		outChannel <- message
	}()
	return outChannel
}

func main() {
	runtime.GOMAXPROCS()
	fanOut := FanOut("message1")

	fmt.Println(<-fanOut)
	fmt.Println(<-fanOut)
}
