package main

import "fmt"

func main() {

	dist := hammingDistance(1, 4)
	fmt.Println(dist)
}

//The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
//XOR gives 0 when same bits else 1
func hammingDistance(x int, y int) int {

	answer := x ^ y
	//count number of set bits
	return CounttheSetBits(answer)

}

func CounttheSetBits(n int) int {
	count := 0
	for n > 0 {

		if (n & 1) == 1 {
			count++
		}
		n = n >> 1

	}
	return count
}
