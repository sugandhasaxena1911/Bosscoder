package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6

	positions := searchRange(nums, target)
	fmt.Println(positions)

}

func searchRange(nums []int, target int) []int {
	positions := []int{}

	positions = append(positions, findFirstOccurence(nums, target))
	positions = append(positions, findLastOccurence(nums, target))
	return positions
}

func findFirstOccurence(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	m := 0
	pos := -1
	for l <= r {
		m = (l + r) / 2
		if target == nums[m] {
			pos = m
			r = m - 1
			continue
		}
		if target < nums[m] {
			r = m - 1
		} else if target > nums[m] {
			l = m + 1
		}

	}
	if pos != -1 {
		return pos
	} else {
		return -1

	}
}

func findLastOccurence(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	m := 0
	pos := -1
	for l <= r {
		m = (l + r) / 2
		if target == nums[m] {
			pos = m
			l = m + 1
			continue
		}
		if target > nums[m] {
			l = m + 1
		} else if target < nums[m] {
			r = m - 1
		}

	}
	if pos != -1 {
		return pos
	} else {
		return -1

	}
}
