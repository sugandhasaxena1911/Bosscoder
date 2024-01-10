package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	res := subsets(nums)
	fmt.Println(res)
}

var results [][]int

func subsets(nums []int) [][]int {
	results = [][]int{}
	SubsetsRec(nums, 0, []int{})
	return results

}

func SubsetsRec(nums []int, index int, curr []int) {
	if len(nums) == index {
		//fmt.Println(index, curr)
		results = append(results, curr)
		return
	}
	//fmt.Println("here ", index, curr)
	// character part of subset & then not part of subset
	n := append([]int{}, curr...)
	n = append(n, nums[index])
	SubsetsRec(nums, index+1, n)

	m := append([]int{}, curr...)
	SubsetsRec(nums, index+1, m)

}
