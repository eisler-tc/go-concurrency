package main

import "fmt"

func FanIn(message1, message2 <-chan string) chan string {
	outChannel := make(chan string)
	go func() {
		outChannel <- <-message1
	}()
	go func() {
		outChannel <- <-message2
	}()

	return outChannel
}

func main() {

	message1Chn := make(chan string)
	message2Chn := make(chan string)
	fanInChannel := FanIn(message1Chn, message2Chn)
	message1Chn <- "message1"
	message2Chn <- "message2"

	fmt.Println(<-fanInChannel)
	fmt.Println(<-fanInChannel)
}
