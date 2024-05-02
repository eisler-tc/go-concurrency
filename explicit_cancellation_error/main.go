package main

import (
	"fmt"
	"sync"
	"time"
)

var fakeDataSet = []string{"data1", "data2", "data3", "data4", "data5", "data6", "data7", "data8", "data9", "data10"}

func GetDataFromExternalDataSource(wg *sync.WaitGroup, retrievedData chan<- string, done <-chan interface{}) {
	defer wg.Done()
	for _, fakeData := range fakeDataSet {
		wg.Add(1)
		select {
		case <-done:
			fmt.Println("Exiting from GetDataFromExternalDataSource")
			return
		default:
			retrievedData <- fakeData
			wg.Done()
		}
	}
}

func SendDataToExternalDataSource(wg *sync.WaitGroup, processedData chan string, done <-chan interface{}) {
	defer wg.Done()
	for {
		wg.Add(1)
		select {
		case <-done:
			fmt.Println("Exiting from SendDataToExternalDataSource")
			return
		case data := <-processedData:
			fmt.Println("sending to an external data source |" + data + "|")
			wg.Done()
		}

	}
}

func ProcessData(wg *sync.WaitGroup, retrievedData chan string, processedData chan string, done <-chan interface{}) {
	defer wg.Done()
	for {
		wg.Add(1)
		select {
		case <-done:
			fmt.Println("Exiting from ProcessData")
			return
		case data := <-retrievedData:
			time.Sleep(time.Second)
			processedData <- data + "-processed"
			wg.Done()
		}
	}
}

func main() {

	wg := sync.WaitGroup{}

	retrievedData := make(chan string)
	processedData := make(chan string)
	done := make(chan interface{}, 3)

	go GetDataFromExternalDataSource(&wg, retrievedData, done)
	go ProcessData(&wg, retrievedData, processedData, done)
	go SendDataToExternalDataSource(&wg, processedData, done)

	time.Sleep(3 * time.Second)
	fmt.Println("An explicit cancellation signal came..")

	//send done signal to all af 3 routines
	done <- struct{}{}
	done <- struct{}{}
	done <- struct{}{}

	wg.Wait()
}
