package main

import (
	"fmt"
)

// It’s possible to read data from a already closed buffered channel.
// The channel will return the data that is already written to the channel and once all the data has been read,
// it will return the zero value of the channel.
func main() {
	ch := make(chan int, 5)
	ch <- 5
	ch <- 6
	close(ch)
	n, open := <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received aagain : %d, open: %t\n", n, open)
	// however the range loop exits as soon as channel no data to read & channel is closed
	ch = make(chan int, 5)

	ch <- 5
	ch <- 6
	ch <- 7
	close(ch)
	for n := range ch {
		fmt.Println("Received range loop:", n)
	}
}
