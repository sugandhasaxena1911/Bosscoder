package main

import "fmt"

func main() {

	nums := []int{5, 4, 8, 3, 1}
	merged := MergeSort(nums)
	fmt.Println(merged)

}

// O(nlogn)
func MergeSort(nums []int) []int {
	l := 0
	r := len(nums) - 1
	if l >= r {
		return nums
	}

	m := (l + r) / 2
	nums1 := MergeSort(nums[:m+1])
	nums2 := MergeSort(nums[m+1:])
	return Merge(nums1, nums2)
}

func Merge(nums1 []int, nums2 []int) []int {
	pt1 := 0
	pt2 := 0
	merged := []int{}
	for pt1 < len(nums1) && pt2 < len(nums2) {

		if nums1[pt1] < nums2[pt2] {
			merged = append(merged, nums1[pt1])
			pt1++
		} else {
			merged = append(merged, nums2[pt2])
			pt2++
		}
	}

	for pt1 < len(nums1) {
		merged = append(merged, nums1[pt1])
		pt1++

	}

	for pt2 < len(nums2) {
		merged = append(merged, nums2[pt2])
		pt2++
	}

	return merged
}
