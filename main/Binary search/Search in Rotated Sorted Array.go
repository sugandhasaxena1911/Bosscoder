package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	k := search(nums, target)
	fmt.Println(k)
}

func findMinIndex(nums []int) int {
	n := len(nums)
	low, high := 0, n-1
	for low < high {
		mid := (low + high) / 2
		// nums[mid] > nums[high]  means
		//1)  mid is not the smallest element ..
		//2)  first half is sorted & pivot lies in the second half  , therefore low = mid +1
		/*if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}*/
		if nums[mid] < nums[high] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low /// any of low or high can be returned , since l==r

}

func Binarysearch(nums []int, target int, l int, r int) int {
	//l := 0
	//r := len(nums) - 1
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

func search(nums []int, target int) int {
	pivot := findMinIndex(nums)
	fmt.Println("Pivot or min element index ", pivot)
	// sorted array in ascending
	if pivot == 0 {
		return Binarysearch(nums, target, 0, len(nums)-1)

	}
	if target >= nums[0] {
		return Binarysearch(nums, target, 0, pivot-1)
	} else {
		return Binarysearch(nums, target, pivot, len(nums)-1)
	}
}
