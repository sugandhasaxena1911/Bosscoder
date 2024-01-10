package main

import "fmt"

// from the notes, Maximum value of |arr[i] – arr[j]| + |i – j|
func main() {
	nums := []int{1, 2, 3, 1}
	max := MaximumValue(nums)
	fmt.Println("MAX ", max)

	nums = []int{1, 2, 3, 1}
	max = MaximumValueV2(nums)
	fmt.Println("MAX-2nd algo  ", max)
}

// time O(n2)
func MaximumValue(nums []int) int {
	maximum := nums[0] - nums[1] + 1 - 0
	// brute

	for j := 0; j < len(nums); j++ {
		for i := j + 1; i < len(nums); i++ {
			maximum = max(maximum, (nums[i]-nums[j])+(i-j))
			fmt.Println(maximum)
		}
	}
	return maximum

}

// time :O(n), extra space O(n)
func MaximumValueV2(nums []int) int {
	// brute
	arr1 := make([]int, len(nums))
	arr2 := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		arr1[i] = nums[i] + i
		arr2[i] = nums[i] - i
	}
	max1 := arr1[0]
	min1 := arr1[0]
	for x := 1; x < len(arr1); x++ {
		max1 = max(max1, arr1[x])
		min1 = min(min1, arr1[x])
	}

	max2 := arr2[0]
	min2 := arr2[0]
	for x := 1; x < len(arr2); x++ {
		max2 = max(max2, arr2[x])
		min2 = min(min2, arr2[x])
	}

	return max((max1 - min1), (max2 - min2))

}

func max(n1 int, n2 int) int {

	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func min(n1 int, n2 int) int {

	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}
