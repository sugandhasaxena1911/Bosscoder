package main

import (
	"fmt"
	"sync"
)

// 2 routines sending values to both channels
// here we have to use wiat groups. WHy ? wait till all values have been send  then close channel
// semaphores/channels  can be used to Accomplish this without using wait groups
// check code Semaphore.go
var wg sync.WaitGroup

func main() {
	c := make(chan int)
	wg.Add(2) // this should be in the main function, if we add both wg.Add in separate go routines then race condition : because
	go GoRoutine1(c)
	go GoRoutine2(c)
	go func() {
		for x := range c {
			fmt.Println(x)
		}
	}()
	// above code can result in deadlock & we need to close channel ..However you get error that send to closed channels
	// so close channels once all values have been send.  wait till all values are send
	wg.Wait()
	go GoCloseChannel(c)

	fmt.Println("end main")

}

func GoRoutine1(c chan int) {

	for x := 1; x <= 100; x++ {
		c <- x
	}
	wg.Done()
}

func GoRoutine2(c chan int) {
	for x := 101; x <= 200; x++ {
		c <- x
	}
	wg.Done()
}

func GoCloseChannel(c chan int) {
	close(c)
}
