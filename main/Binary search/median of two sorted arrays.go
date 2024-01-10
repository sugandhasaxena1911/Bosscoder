package main

import "fmt"

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	median := findMedianSortedArrays(nums1, nums2)
	fmt.Println(median)

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr3 := []int{}
	median := 0.0
	x := 0
	y := 0
	for x < len(nums1) && y < len(nums2) {
		if nums1[x] < nums2[y] {
			arr3 = append(arr3, nums1[x])
			x++
		} else {
			arr3 = append(arr3, nums2[y])
			y++
		}
	}
	fmt.Println(arr3)

	for x < len(nums1) {
		fmt.Println("here")
		arr3 = append(arr3, nums1[x])
		x++

	}
	for y < len(nums2) {
		fmt.Println("here2")

		arr3 = append(arr3, nums2[y])
		y++
	}
	fmt.Println(arr3)
	low := 0
	high := len(arr3) - 1
	m := (low + high/2)

	if len(arr3)%2 == 0 {
		fmt.Println(x)
		median = (float64(arr3[m]+arr3[m+1]) / 2.0)
	} else {
		median = float64(arr3[m])
	}
	return median
}
