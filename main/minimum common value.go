package main

import "fmt"

func main() {

	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4}
	c := getMinCommon(nums1, nums2)
	fmt.Println(c)
}

func getMinCommon(nums1 []int, nums2 []int) int {
	// arrays are sorted
	i := 0
	j := 0
	common := -1

	for i < len(nums1) && j < len(nums2) {

		if nums1[i] < nums2[j] {
			i++
			continue
		}
		if nums1[i] > nums2[j] {
			j++
			continue
		}
		// values are equal
		fmt.Println("common found ", nums1[i])
		common = nums1[i]
		break
	}

	return common

}

func getMinCommon2(nums1 []int, nums2 []int) int {
	// arrays are sorted
	i := 0
	j := 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			return nums1[i]
		}

		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		}

	}

	return -1

}
