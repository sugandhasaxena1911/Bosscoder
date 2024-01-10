package main

import "fmt"

func main() {
	nums := []int{1, 5, 6, 9}
	fmt.Println(nums)
	BubbleSort(nums)
	fmt.Println(nums)
}

//At each  iteration moves largest at the end
//worst O(n2), best O(n)- find if array is alreday sorted by swapped indicator
func BubbleSort(nums []int) {
	high := len(nums) - 1
	low := 0
	swapped := false
	for low < high {
		fmt.Println("inside outer")
		for x := low + 1; x <= high; x++ {
			fmt.Println("inside inner")

			if nums[x] < nums[x-1] {
				nums[x-1], nums[x] = nums[x], nums[x-1]
				swapped = true
			}
		}
		if swapped == false {
			break
		}
		high--
	}

}
