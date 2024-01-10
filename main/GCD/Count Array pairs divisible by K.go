package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := 2

	count := countPairs(nums, k)
	fmt.Println(count)

}

// O(n*2)
func countPairs(nums []int, k int) int64 {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i]*nums[j])%k == 0 {
				count++
			}
		}
	}

	return int64(count)
}
