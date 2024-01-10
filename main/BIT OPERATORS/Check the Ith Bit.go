package main

import "fmt"

func main() {
	n := 137
	k := 5
	result := FindKthBit(n, k)
	fmt.Printf("the binary of number %d is %b\n", n, n)
	fmt.Println("the kth bit : ", k, " is ", result)

}

// true : 1 , false :0
func FindKthBit(n int, k int) int {

	if (n & (1 << k)) == 0 {
		return 0
	} else {
		return 1

	}

}
