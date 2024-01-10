package main

import (
	"fmt"
	"sort"
)

func main() {
	target := 2
	nums := []int{-1, 1, 2, 3, 1}

	result := countPairs(nums, target)
	fmt.Println(result)
}

func countPairs(nums []int, target int) int {
	n := 0
	for x := 0; x < len(nums); x++ {
		for y := x + 1; y < len(nums); y++ {
			if nums[x]+nums[y] < target {
				n++
			}

		}
	}

	return n
}

func countPairs2(nums []int, target int) int {
	sort.Ints(nums)

	ans, left, right := 0, 0, len(nums)-1
	for left < right {
		if nums[right]+nums[left] < target {
			ans += right - left
			left++
		} else {
			right--
		}
	}
	return ans
}
