package main

import (
	"bufio"
	"fmt"
	"github.com/goinggo/workpool"
	"os"
	"runtime"
	"time"
)

type MyWork struct {
	WP *workpool.WorkPool
	Fn func()
}

func (workPool *MyWork) DoWork(workRoutine int) {
	workPool.Fn()
	//panic("test")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	workPool := workpool.New(2, 2)

	shutdown := false // Just for testing, I Know

	go func() {
		work1 := &MyWork{WP: workPool, Fn: func() {
			for {
				time.Sleep(time.Second)
				fmt.Println("I did some expensive job..")
			}

		}}
		work2 := &MyWork{WP: workPool, Fn: func() {
			for {
				fmt.Println("I did some cheap job..")
			}
		}}

		err := workPool.PostWork("name_routine1", work1)
		err = workPool.PostWork("name_routine2", work2)

		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			time.Sleep(100 * time.Millisecond)
		}
		if shutdown == true {
			return
		}
	}()
	fmt.Println("Press any key to finish execution...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	shutdown = true

	fmt.Println("Shutting Down\n")
	workPool.Shutdown("name_routine")
}
