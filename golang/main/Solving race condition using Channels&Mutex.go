package main

import (
	"fmt"
	"sync"
)

// In general use channels when Goroutines need to communicate with each other and mutexes when
// only one Goroutine should access the critical section of code.
var wg sync.WaitGroup
var mu sync.Mutex
var x int

func main() {
	//ch := make(chan int, 1)
	for y := 1; y <= 500; y++ {
		wg.Add(1)
		// here race condition will occur when go routines can fetch wrong values of x
		//& final value will be different
		//go incrementX()

		//go incrementX_CH(ch)
		go incrementX_MUTEX()
	}
	wg.Wait()
	fmt.Println("final value is ", x)

}

func incrementX() {
	x++
	wg.Done()
}

func incrementX_CH(ch chan int) {
	ch <- 1 // this line wont be executed & blocked until another go routine has written value of x & then read from buffered channel
	x++
	<-ch
	wg.Done()
}

func incrementX_MUTEX() {
	mu.Lock()
	x++
	mu.Unlock()
	wg.Done()
}
