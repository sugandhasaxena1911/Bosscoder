package main

import "fmt"

func main() {
	n := 137
	k := 3
	result := ToggleKthBit(n, k)
	fmt.Printf("the binary of number %d is %b\n", n, n)
	fmt.Printf("the binary of result %d is %b", result, result)

}

func ToggleKthBit(n int, k int) int {

	return n ^ (1 << k)

}
