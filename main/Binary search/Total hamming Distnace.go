package main

import "fmt"

func main() {

	nums := []int{4, 14, 4}
	dist := totalHammingDistance(nums)
	fmt.Println(dist)
}

//PS : there is a better way that i need to see

//Given an integer array nums,
//return the sum of Hamming distances between all the pairs of the integers in nums.
func totalHammingDistance(nums []int) int {
	s := 0
	for x := 0; x < len(nums)-1; x++ {
		for y := x + 1; y < len(nums); y++ {
			fmt.Println(nums[x], nums[y])
			d := hammingDistance(nums[x], nums[y])
			fmt.Println("Hamming Distance d ", d)
			s = s + d
		}
	}
	return s
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
