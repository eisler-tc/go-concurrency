package main

import (
	"fmt"
	"sync"
	"time"
)

func SumCheck(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func sum(wg *sync.WaitGroup, s []int, sumChan chan int) {
	defer wg.Done()
	sum := 0
	for _, v := range s {
		time.Sleep(time.Second)
		sum += v
	}
	sumChan <- sum
}

func main() {
	start := time.Now()

	wg := sync.WaitGroup{}

	sampleList := make([]int, 10)
	for i := 0; i < 10; i++ {
		sampleList[i] = i
	}

	sumChan := make(chan int)

	wg.Add(2)
	go sum(&wg, sampleList[:len(sampleList)/2], sumChan)
	go sum(&wg, sampleList[len(sampleList)/2:], sumChan)

	halfSum1, halfSum2 := <-sumChan, <-sumChan

	total := halfSum1 + halfSum2
	wg.Wait()
	correctSum := SumCheck(sampleList)
	if correctSum != total {
		fmt.Printf("expected global sum was %d, but actual: %d\n", correctSum, total)
	}
	fmt.Printf("total sum=%d\n", total)
	fmt.Printf("Total execution time : %v", time.Since(start))
}
