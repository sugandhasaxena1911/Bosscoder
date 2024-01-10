package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{1, 2, 3, 4, 5}
	k := 3
	score := maximizeSum(nums, k)
	fmt.Println(score)
	nums = []int{1, 2, 3, 4, 5}

	score2 := maximizeSum2(nums, k)
	fmt.Println(score2)

}
func maximizeSum(nums []int, k int) int {
	sort.Ints(nums)
	l := len(nums)
	score := 0
	m := 0
	for k > 0 {
		m = nums[l-1]
		fmt.Println(m)
		score = score + m
		nums[l-1] = m + 1
		k--
	}
	return score
}
func maximizeSum2(nums []int, k int) int {
	sort.Ints(nums)
	l := len(nums)
	score := 0
	m := nums[l-1]
	score = k*m + ((k-1)*k)/2

	return score
}

//k=3
//m+0   m+1   m+2
