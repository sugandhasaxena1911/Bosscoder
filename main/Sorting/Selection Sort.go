package main

import "fmt"

func main() {
	nums := []int{8, 9, 10, 11}
	fmt.Println(nums)
	selectionSort(nums)
	fmt.Println(nums)
}

//at each iteration finds min from the right array & moves min at the start
//O(n2)  , best case O(n2)-- even if array sorted the loop will run
func selectionSort(nums []int) {
	low := 0
	high := len(nums) - 1
	min := low
	for low < high {
		fmt.Println("inside outer")
		for x := low + 1; x <= high; x++ {
			fmt.Println("inside inner")

			if nums[x] < nums[min] {
				min = x
			}
		}
		//swap smallest and nums[low]
		nums[low], nums[min] = nums[min], nums[low]
		low++
		min = low
	}
}
