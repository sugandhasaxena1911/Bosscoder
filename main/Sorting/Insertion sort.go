package main

import "fmt"

func main() {
	nums := []int{0, 3, 5}
	InsertionSort(nums)
	fmt.Println(nums)
}

//left side is sorted at each iteration & at each iteration tries to insert key at correct place
// O(n2), best O(n)
func InsertionSort(nums []int) {
	low := 0
	high := len(nums) - 1

	for low < high {
		fmt.Println("Outer")
		key := nums[low+1]
		x := low
		for x >= 0 && nums[x] > key {
			fmt.Println("Inner")
			nums[x+1] = nums[x]
			x--
		}
		nums[x+1] = key
		low++
	}

}
