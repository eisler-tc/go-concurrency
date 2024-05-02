package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("wrote", i, "to ch")
	}
	close(ch)
}

func main() {
	ch := make(chan int, 5)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read ", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}
