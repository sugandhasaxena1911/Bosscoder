package main

import "fmt"

func main() {
	arr := []int{1, 2, 1}
	final := getConcatenation(arr)
	fmt.Println(final)
	fmt.Println(getConcatenation2(arr))
	fmt.Println(getConcatenation3(arr))

}
func getConcatenation(nums []int) []int {
	var answer []int
	answer = append(answer, nums...)
	answer = append(answer, nums...)
	return answer
}

// thsi method is more better than 1st or last approach
func getConcatenation2(nums []int) []int {
	n := len(nums)
	answer := make([]int, n*2)
	for x := 0; x < n*2; x++ {
		answer[x] = nums[x%n]

	}
	return answer
}

func getConcatenation3(nums []int) []int {

	return append(nums, nums...)
}
