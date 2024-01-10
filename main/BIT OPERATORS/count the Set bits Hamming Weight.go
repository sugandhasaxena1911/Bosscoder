package main

import "fmt"

func main() {
	n := 21
	k := CounttheSetBits(n)
	fmt.Printf("the total number of set bits in %b is %d ", n, k)

	n = 21
	k = CounttheSetBitsV2(n)
	fmt.Printf("\nthe total number of set bits in %b is %d ", n, k)
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

func CounttheSetBitsV2(n int) int {
	count := 0
	k := 0
	for k <= 31 {

		if (n & (1 << k)) != 0 {
			count++
		}
		k++

	}
	return count
}
