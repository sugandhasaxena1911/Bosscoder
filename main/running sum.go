package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	res := runningSum(nums)
	fmt.Println(res)

}

func runningSum(nums []int) []int {

	for x := 1; x < len(nums); x++ {
		nums[x] = nums[x] + nums[x-1]
	}

	return nums

}
