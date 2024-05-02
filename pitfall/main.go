package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var a = 0
var b = 0
var c = 0

func doSomeExpensiveJob(ctx context.Context, wg *sync.WaitGroup) {
	for {
		wg.Add(1)
		select {
		case <-ctx.Done():
			fmt.Println("Exiting from doSomeExpensiveJob method")
			wg.Done()
			return
		default:
			runtime.LockOSThread()
			time.Sleep(time.Second)
			fmt.Println("I did some expensive job..")
			a++
			wg.Done()
		}
	}

}

func doSomeOtherExpensiveJob(ctx context.Context, wg *sync.WaitGroup) {
	for {
		wg.Add(1)
		select {
		case <-ctx.Done():
			fmt.Println("Exiting from doSomeOtherExpensiveJob method")
			wg.Done()
			return
		default:
			runtime.LockOSThread()
			time.Sleep(time.Second)
			fmt.Println("I did some  other expensive job..")
			b++
			wg.Done()
		}
	}

}

func doSomeCheapJob(ctx context.Context, wg *sync.WaitGroup) {
	for {
		wg.Add(1)
		select {
		case <-ctx.Done():
			fmt.Println("Exiting from doSomeCheapJob method")
			wg.Done()
			return
		default:
			//fmt.Println("I did some cheap job..")
			c++
			wg.Done()
		}
	}

}

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	wg := sync.WaitGroup{}

	go doSomeExpensiveJob(ctx, &wg)
	go doSomeOtherExpensiveJob(ctx, &wg)
	go doSomeCheapJob(ctx, &wg)

	<-ctx.Done()
	wg.Wait()
	fmt.Printf("a=%v", a)
	fmt.Printf("b=%v", b)
	fmt.Printf("c=%v", c)
}
