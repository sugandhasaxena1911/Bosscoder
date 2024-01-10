package main

import "fmt"

func main() {
	nums := []int{4, 5, 1, 2, 3}

	min := findMin(nums)
	fmt.Println(min)

}

// in this approach we dont check mid element equality
//if we do , then we need to check if mid == di element , then mid-1 & mid+1 comes to picture
// then out o bound issue .. previous=(mid + N - 1) % N and next = (mid + 1) % N     --> COMPLEX CODE
// theeefore below approach
func findMin(nums []int) int {
	n := len(nums)
	low, high := 0, n-1
	for low < high {
		mid := (low + high) / 2
		// nums[mid] > nums[high]  means
		//1)  mid is not the smallest element ..
		//2)  first half is sorted & pivot lies in the second half  , therefore low = mid +1
		/*if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}*/
		if nums[mid] < nums[high] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return nums[low] /// any of low or high can be returned , since l==r

}
