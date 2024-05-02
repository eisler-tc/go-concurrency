package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var fakeDataSet = []string{"data1", "data2", "data3", "data4", "data5", "data6", "data7", "data8", "data9", "data10"}

func GetDataFromExternalDataSource(ctx context.Context, wg *sync.WaitGroup, retrievedData chan string) {
	defer wg.Done()
	for _, fakeData := range fakeDataSet {
		select {
		case <-ctx.Done():
			fmt.Println("Context time out is exceeded. Exiting from GetDataFromExternalDataSource")
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

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	wg.Add(1)
	go GetDataFromExternalDataSource(ctx, &wg, retrievedData)
	go ProcessData(&wg, retrievedData, processedData)
	go SendDataToExternalDataSource(&wg, processedData)

	wg.Wait()
}
