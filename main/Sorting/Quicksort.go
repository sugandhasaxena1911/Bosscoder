package main

import "fmt"

func main() {
	nums := []int{4, 5, 8, 9, 11}
	QuickSort(nums)

}

//time complexity O(nlogn), worst O(n2)
// quicksort : when we have random access support like in arrays
//merge sort : when we dont have randon acces supoort like in Linked List
func QuickSort(nums []int) {

	QuickSortRec(nums, 0, len(nums)-1)
	fmt.Println("After final", nums)

}

func QuickSortRec(nums []int, low int, high int) {
	if low < high {
		pivot := FindPivotIndex(nums, low, high)
		fmt.Println("After finding pivot", nums, pivot)
		fmt.Println("callling for ", low, pivot-1)

		QuickSortRec(nums, low, pivot-1)
		fmt.Println("After left half", nums, pivot)
		fmt.Println("callling for ", pivot+1, high)

		QuickSortRec(nums, pivot+1, high)
		fmt.Println("After right half", nums, pivot)

	}

}
func FindPivotIndex(nums []int, low int, high int) int {

	pivot := low
	for low <= high {
		if nums[low] <= nums[pivot] {
			low++
			continue
		}

		if nums[high] > nums[pivot] {
			high--
			continue

		}
		//swap low and high
		nums[low], nums[high] = nums[high], nums[low]

	}
	// here low>high, assign pivot value to correct index
	nums[high], nums[pivot] = nums[pivot], nums[high]

	// return correct pivot index
	return high
}
