package main

import "fmt"

func main() {
	ch := make(chan int)
	done := make(chan int)
	go Routine1(ch)
	go Read1(ch, done)
	go Read2(ch, done)
	<-done
	<-done
	fmt.Println("END")

}

func Routine1(ch chan int) {
	for x := 1; x <= 20; x++ {
		ch <- x
	}
	close(ch)
}

func Read1(ch chan int, done chan int) {
	for x := range ch {
		fmt.Println("Read1 ", x)
	}
	done <- 1
}
func Read2(ch chan int, done chan int) {
	for x := range ch {
		fmt.Println("Read2", x)
	}
	done <- 1
}
