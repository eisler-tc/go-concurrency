package main

import "time"

func main() {

	waitForMyWork := make(chan struct{})
	myWorkIsCompleted := make(chan struct{})
	//anotherRoutineWork := make(chan struct{})

	//go routineWork(anotherRoutineWork)
	go complexJob(waitForMyWork, myWorkIsCompleted)

	for {
		select {
		case <-waitForMyWork:
			println("okay I am waiting. for complex job")
			<-myWorkIsCompleted
			println("complex job is finished")
		default:
			println("I am doing regular work")
			time.Sleep(6 * time.Second)
			println("I am done with my regular work")
		}
	}
}

//
//func routineWork(anotherRoutineWork chan struct{}) {
//	for {
//		anotherRoutineWork <- struct{}{}
//		time.Sleep(10 * time.Second)
//	}
//}

func complexJob(waitForMyWork chan struct{}, myWorkIsCompleted chan struct{}) {
	for {
		time.Sleep(3 * time.Second)
		waitForMyWork <- struct{}{}
		time.Sleep(2 * time.Second)
		myWorkIsCompleted <- struct{}{}
	}
}
