package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2}
	r := permute(nums)
	fmt.Println(r)

}

var results [][]int

func permute(nums []int) [][]int {
	sort.Ints(nums) // IMPORTANT
	permutationArray2(nums, []int{})
	//fmt.Println(results)
	return results
}

func permutationArray2(nums []int, ans []int) {
	if len(nums) == 0 {
		fmt.Println("HELLO ", ans)

		results = append(results, ans)
		return
	}
	fmt.Println(nums)
	for x := 0; x < len(nums); x++ {
		if x-1 >= 0 {
			if nums[x] == nums[x-1] {
				continue
			}
		}
		arr := append([]int{}, nums[0:x]...)
		arr = append(arr, nums[x+1:]...)
		fmt.Println(nums, arr)
		n := append([]int{}, ans...)
		n = append(n, nums[x])
		fmt.Println("arr ", arr, n, ans, x, nums[x])
		permutationArray2(arr, n)

	}

}
