package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	qnums := [][]int{{1, 3}, {2, 4}}
	sum := RangeSumQueries(nums, qnums)
	fmt.Println("Total ", sum)

	nums = []int{1, 2, 3, 4, 5}
	qnums = [][]int{{1, 3}, {2, 4}}
	sum = RangeSumQueriesV2(nums, qnums)
	fmt.Println("Total ", sum)

}

//BRUTE FORCE
// Complexity : O(q*n)  if q=n then O(n*2)
func RangeSumQueries(nums []int, qnums [][]int) int {
	q := len(qnums)
	sum := 0

	for x := 0; x < q; x++ {
		left := qnums[x][0]
		right := qnums[x][1]

		// subarray left index & right index calculated
		//parse over 1Darray
		s := 0

		for i := left; i <= right; i++ {
			s = s + nums[i]

		}
		fmt.Println(s)
		sum = sum + s
	}
	return sum

}

// Using PRE-PROCESSING Time O(q)+O(n)=O(n+q)  if n==q   O(2n)=O(n), extra space : O(n)
func RangeSumQueriesV2(nums []int, qnums [][]int) int {

	presum := make([]int, len(nums))
	presum[0] = nums[0]
	// O(n) to cal presum array
	for i := 1; i < len(nums); i++ {
		presum[i] = nums[i] + presum[i-1]
	}

	q := len(qnums)
	sum := 0
	for x := 0; x < q; x++ {
		left := qnums[x][0]
		right := qnums[x][1]
		s := presum[right] - presum[left-1] // O(1) time complexity
		sum = sum + s
		fmt.Println(s)

	}
	return sum
}
