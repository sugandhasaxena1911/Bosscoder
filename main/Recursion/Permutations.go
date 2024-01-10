package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	r := permute(nums)
	fmt.Println(r)

}

var results [][]int

func permute(nums []int) [][]int {
	permutationArray(nums, []int{})
	//fmt.Println(results)
	return results
}

func permutationArray(nums []int, ans []int) {
	if len(nums) == 0 {
		fmt.Println("HELLO ", ans)
		results = append(results, ans)
	}
	fmt.Println(nums)

	for x := 0; x < len(nums); x++ {
		arr := append([]int{}, nums[0:x]...)
		arr = append(arr, nums[x+1:]...)
		fmt.Println(nums, arr)
		n := append([]int{}, ans...)
		n = append(n, nums[x])
		fmt.Println("arr ", arr, n, ans, x, nums[x])
		permutationArray(arr, n)

	}

}
