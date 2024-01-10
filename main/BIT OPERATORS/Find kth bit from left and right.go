package main

import "fmt"

func main() {
	n := 137
	k := 7
	left := FindKthBitLeft(n, k)
	right := FindKthBitRight(n, k)
	fmt.Printf("Number %b\n", n)
	fmt.Printf("the Kth %d left bit is %d & right bit is %d", k, left, right)

}

func FindKthBitLeft(n int, k int) int {
	if (n & (1 << (31 - k))) != 0 {
		return 1
	} else {
		return 0
	}
}

func FindKthBitRight(n int, k int) int {

	if (n & (1 << k)) != 0 {
		return 1
	} else {
		return 0
	}

}
