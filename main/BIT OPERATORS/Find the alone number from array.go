package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 5, 1, 3, 2, 5, 6}
	n := FindLoneNumber(nums)
	fmt.Println(n)

}

func FindLoneNumber(nums []int) int {
	result := nums[0]

	for x := 1; x < len(nums); x++ {
		result = result ^ nums[x]
	}
	return result
}
