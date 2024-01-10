package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 3, 1}
	nextpermutation(nums)

}

func nextpermutation(nums []int) {
	check := 0

	for x := len(nums) - 1; x > 0; x-- {
		fmt.Println("inside x = ", x, " subaaray ", nums[x:])
		nums2 := nums[x:] //subarray
		//element := nums[x-1]

		// find the smallest elelemnt of subaaray
		n := 10
		idx := -1
		for i, _ := range nums2 {
			if nums2[i] > nums[x-1] {

				n = min(n, nums2[i])
				idx = i
			}

		}
		fmt.Println("Next bigger element in ", nums2, " is ", n, " compare with ", nums[x-1])
		if n > nums[x-1] && idx != -1 {
			fmt.Println("found")
			nums[x-1], nums2[idx] = n, nums[x-1]
			fmt.Println(nums[0:x], nums2)

			sort.Ints(nums2)
			nums = append(nums[0:x], nums2...)
			check = 1
			break
		}

	}
	if check == 0 {
		sort.Ints(nums)
	}
	fmt.Println(nums)
}

func min(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}
