package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go func() {

		for x := 0; x <= 15; x++ {
			ch <- x
		}
	}()

	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()
	time.Sleep(5 * time.Second)
}
