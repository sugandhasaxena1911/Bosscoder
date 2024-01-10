package main

import "fmt"

func main() {
	fmt.Println("I am in main go routine")
	//lets call squares & cubes in different go routines
	chsq := make(chan int)
	chcube := make(chan int)

	go CalcCube(5, chcube)
	go CalcSquares(5, chsq)
	fmt.Println("the cube & square is ", <-chcube, <-chsq)

}

func CalcCube(n int, chcube chan<- int) {
	sum := 0
	for n > 0 {
		lg := n % 10
		sum = sum + (lg * lg * lg)
		n = n / 10
	}
	fmt.Println("the cube is ", sum)
	chcube <- sum

}

func CalcSquares(n int, chsq chan<- int) {
	sum := 0
	for n > 0 {
		lg := n % 10
		sum = sum + (lg * lg)
		n = n / 10
	}
	fmt.Println("the sqaures is ", sum)
	chsq <- sum
}
