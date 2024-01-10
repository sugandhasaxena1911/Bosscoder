package main

import (
	"fmt"
)

func main() {
	ans := smallerNumbersThanCurrent([]int{1, 2, 3})
	fmt.Println(ans)
}

func smallerNumbersThanCurrent(nums []int) []int {

	result := make([]int, len(nums))
	cnt := 0
	for x := 0; x < len(nums); x++ {
		cnt = 0

		for y := x + 1; y < len(nums); y++ {
			if nums[y] < nums[x] {
				cnt++
			}
		}

		for z := 0; z < x; z++ {
			if nums[z] < nums[x] {
				cnt++
			}
		}

		result[x] = cnt

	}

	return result
}
