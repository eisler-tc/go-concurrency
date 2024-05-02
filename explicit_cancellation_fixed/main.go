package main

import (
	"fmt"
	"sync"
	"time"
)

var fakeDataSet = []string{"data1", "data2", "data3", "data4", "data5", "data6", "data7", "data8", "data9", "data10"}

func GetDataFromExternalDataSource(wg *sync.WaitGroup, retrievedData chan string, done <-chan interface{}) {
	defer wg.Done()
	for _, fakeData := range fakeDataSet {
		select {
		case <-done:
			fmt.Println("Exiting from GetDataFromExternalDataSource")
			close(retrievedData)
			return
		default:
			retrievedData <- fakeData
			fmt.Println("pushed " + fakeData)
			wg.Add(2)
		}
	}
}

func SendDataToExternalDataSource(wg *sync.WaitGroup, processedData <-chan string) {
	for {
		data, ok := <-processedData
		if !ok {
			fmt.Println("no data from processedData channel")
			return
		}
		fmt.Println("sending to an external data source |" + data + "|")
		wg.Done()
	}
}

func ProcessData(wg *sync.WaitGroup, retrievedData <-chan string, processedData chan<- string) {
	defer fmt.Println("Exiting from ProcessData")
	for {
		data, ok := <-retrievedData
		if !ok {
			fmt.Println("no data from retrievedData channel")
			return
		}
		time.Sleep(time.Second)
		processedData <- data + "-processed"
		wg.Done()
	}
}

func main() {

	wg := sync.WaitGroup{}

	retrievedData := make(chan string)
	processedData := make(chan string)
	done := make(chan interface{}, 3)

	wg.Add(1)
	go GetDataFromExternalDataSource(&wg, retrievedData, done)
	go ProcessData(&wg, retrievedData, processedData)
	go SendDataToExternalDataSource(&wg, processedData)

	time.Sleep(3 * time.Second)
	fmt.Println("An explicit cancellation signal came after 3 seconds..")
	done <- struct{}{}
	wg.Wait()
}
