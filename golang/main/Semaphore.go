package main

import "fmt"

func main() {
	ch := make(chan int)
	done := make(chan int)

	go GoRoutine1(ch, done)
	go GoRoutine2(ch, done)
	go ReadValues(ch)

	// at thsi step dont close channel untill all values passed
	<-done // this is stuck till the time all values are send from one channel
	<-done // this is stuck till the time all values are send from one channel

	go CloseChannel(ch)

}

func GoRoutine1(c chan int, done chan int) {

	for x := 1; x <= 100; x++ {
		c <- x
	}

	done <- 1

}

func GoRoutine2(c chan int, done chan int) {
	for x := 101; x <= 200; x++ {
		c <- x
	}
	done <- 1
}

func ReadValues(ch chan int) {
	for x := range ch {
		fmt.Println(x)
	}
}

func CloseChannel(ch chan int) {
	close(ch)
}
