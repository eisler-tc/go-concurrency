package main

import (
	"fmt"
	"math/rand"

	"time"
)

const MaxConcurrentJobs = 2

func main() {

	waitChan := make(chan struct{}, MaxConcurrentJobs)
	count := 0
	for {
		waitChan <- struct{}{}
		count++
		go func(count int) {
			job(count)
			<-waitChan
		}(count)
	}
}

func job(index int) {
	fmt.Println(index, "begin doing something")
	time.Sleep(time.Duration(rand.Intn(10) * int(time.Second)))
	fmt.Println(index, "done")
}
