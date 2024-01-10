package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 2, 5, 9, 7, 4, 8}
	//result := maxProductDifference(nums)
	//fmt.Println(result)

	fmt.Println(nums)

	result := maxProductDifference2(nums)
	fmt.Println(result)
	r := maxProductDifference2(nums)
	fmt.Println(r)

}

func maxProductDifference(nums []int) int {

	sort.Ints(nums)
	l := len(nums)
	return nums[l-1]*nums[l-2] - nums[0]*nums[1]

}

func maxProductDifference2(nums []int) int {
	w, x, y, z := 0, 0, 0, 0
	//4, 2, 5, 9, 7, 4, 8

	for i, _ := range nums {
		fmt.Println(i, nums[i])
		if nums[i] > w {
			x = w
			w = nums[i]
		} else if nums[i] > x || x == 0 {
			x = nums[i]
		}

		if nums[i] < y || y == 0 {
			z = y
			y = nums[i]
		} else if nums[i] < z || z == 0 {
			z = nums[i]
		}
		fmt.Println(w, x, y, z)
	}

	return w*x - y*z
}
