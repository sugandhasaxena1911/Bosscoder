package main

import (
	"fmt"
)

func main() {

	nums := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(nums))

	nums = []int{5}
	fmt.Println(maxSubArray(nums))

	nums = []int{-3}
	fmt.Println(maxSubArray(nums))

	nums = []int{-2, 1}
	fmt.Println(maxSubArray(nums))

}
func maxSubArray(nums []int) int {

	/*if len(nums) == 1 {
		return nums[0]
	}*/
	maxSoFar, currMax := nums[0], nums[0]

	for _, num := range nums[1:] {
		currMax = max(num, currMax+num)
		maxSoFar = max(maxSoFar, currMax)
	}

	return maxSoFar
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}
func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}
