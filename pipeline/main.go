package main

import (
	"fmt"
	"sync"
	"time"
)

var fakeDataSet = []string{"data1", "data2", "data3", "data4", "data5", "data6", "data7", "data8", "data9", "data10"}

func GetDataFromExternalDataSource(wg *sync.WaitGroup, retrievedData chan<- string) {
	defer close(retrievedData)
	defer wg.Done()
	for _, fakeData := range fakeDataSet {
		retrievedData <- fakeData
	}
}

func SendDataToExternalDataSource(wg *sync.WaitGroup, processedDataChan <-chan string) {
	defer wg.Done()
	for processedData := range processedDataChan {
		fmt.Println("sending to an external data source |" + processedData + "|")
	}
}

func ProcessData(wg *sync.WaitGroup, retrievedDataChan <-chan string, processedDataChan chan<- string) {
	defer wg.Done()
	defer close(processedDataChan)
	for retrievedData := range retrievedDataChan {
		time.Sleep(time.Second)
		processedDataChan <- retrievedData + "-processed"
	}
}

func main() {

	wg := sync.WaitGroup{}

	retrievedData := make(chan string)
	processedData := make(chan string)

	wg.Add(3)
	go GetDataFromExternalDataSource(&wg, retrievedData)
	go ProcessData(&wg, retrievedData, processedData)
	go SendDataToExternalDataSource(&wg, processedData)

	wg.Wait()
}
