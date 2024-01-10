package main

import "fmt"

func main() {

	arr := []int{2, 3, 4, -1, -2, 1, 5, 6, -5, 3, 6, 7}
	res := longestSubarray(arr)
	fmt.Println(res)

}

func longestSubarray(nums []int) int {
	size := 0
	sizeprev := 0
	for x := 0; x < len(nums); x++ {
		if nums[x] >= 0 {
			size++
		} else {
			sizeprev = max(size, sizeprev)
			size = 0
		}
	}

	return max(size, sizeprev)

}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
