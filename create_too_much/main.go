package main

import (
	"fmt"
	"github.com/pkg/profile"
	_ "net/http/pprof"
	"sync"
	"time"
)

func doSomeJobs(wg *sync.WaitGroup) {
	fmt.Println("Starting..")
	time.Sleep(time.Second)
	fmt.Println("Job finished")
	wg.Done()
}

func main() {

	//go func() {
	//	http.ListenAndServe(":1234", nil)
	//}()
	defer profile.Start(profile.CPUProfile).Stop()

	wg := sync.WaitGroup{}

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go doSomeJobs(&wg)
	}

	wg.Wait()
	fmt.Println("Execution is finished...")

}
