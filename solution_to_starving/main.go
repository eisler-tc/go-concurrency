package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Job struct {
	Fn func()
}

func main() {

	wg := sync.WaitGroup{}
	runtime.GOMAXPROCS(2)

	a := 0
	b := 0

	job1 := Job{Fn: func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("I did some expensive job..")
			a++
		}

	},
	}

	job2 := Job{Fn: func() {
		for {
			fmt.Println("I did some cheap job..")
			b++
		}

	},
	}

	jobChannel := make(chan *Job, 2)
	waitChan := make(chan struct{}, 2)

	go func() {
		for {
			job := <-jobChannel
			waitChan <- struct{}{}
			go func() {
				job.Fn()
				<-waitChan
			}()
		}
	}()

	go func() {
		for {
			jobChannel <- &job1
			jobChannel <- &job2
		}
	}()

	wg.Wait()
	time.Sleep(10 * time.Second)
	fmt.Printf("a=%v", a)
	fmt.Printf("b=%v", b)
	fmt.Println("Finishing..")
}
