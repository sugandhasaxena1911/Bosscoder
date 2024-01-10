package main

import "fmt"

func main() {

	target := 6
	nums := []int{-1, 0, 3, 5, 9, 12}
	index := search(nums, target)
	fmt.Println(index)

}

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	m := 0
	for l <= r {
		m = ((l + r) / 2)

		if nums[m] == target {
			return m
		}
		if target > nums[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}
