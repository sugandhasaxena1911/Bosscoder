package main

import "fmt"

func main() {

	nums := []int{}
	single := singleNonDuplicate(nums)
	fmt.Println(single)

}

// 1 1 2 2 4 4 5 6 6 7 7
// 0 1 2 3 4 5 6 7 8 9 10    --index
func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if mid%2 == 0 {
			if nums[mid] == nums[mid+1] {
				left = mid + 1
			} else {
				right = mid
			}
		} else if mid%2 == 1 {
			if nums[mid] == nums[mid-1] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	return nums[left]

}
