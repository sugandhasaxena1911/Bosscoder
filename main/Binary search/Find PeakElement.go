package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	index := findPeakElementOptimized(nums)
	fmt.Println(index)

}

//O(logn)
func findPeakElementOptimized(nums []int) int {
	// ascending & descending slopes concepts

	low := 0
	high := len(nums) - 1
	for low < high {

		mid := (low + high) / 2
		//ie part of falling slope, therefre peak lies on the left of mid in ascendign slope
		if nums[mid+1] < nums[mid] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

//brute force  O(n)
func findPeakElement(nums []int) int {
	//peak := 0
	if len(nums) == 1 {
		return 0
	}
	if nums[1] < nums[0] && len(nums) > 1 {
		return 0
	}
	for x := 1; x < len(nums)-1; x++ {
		if nums[x-1] < nums[x] && nums[x+1] < nums[x] {

			return x
		}
	}
	if nums[len(nums)-2] < nums[len(nums)-1] {
		return len(nums) - 1
	}
	return -1
}
