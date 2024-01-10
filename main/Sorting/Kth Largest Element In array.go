package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	kth := findKthLargest(nums, k)
	fmt.Println(kth)

	nums = []int{3, 2, 1, 5, 6, 4}
	k = 2
	kth = findKthLargestCountSort(nums, k)
	fmt.Println(kth)

}

// this will fail if the length of the array is more ..
func findKthLargest(nums []int, k int) int {
	for x := 0; x < k; x++ {

		for y := 1; y < len(nums)-x; y++ {
			if nums[y-1] > nums[y] {
				//swap
				nums[y-1], nums[y] = nums[y], nums[y-1]
			}
		}
	}
	return nums[len(nums)-k]
}

// we can use counr sort since each number -10^4<=nums[i]<=10^4
func findKthLargestCountSort(nums []int, k int) int {
	//largest element could be 10^4 . 10000+10000+1 = length of this array
	// else find max of array in O(n)& then assign the size of arr
	arr := make([]int, 20001)
	for x := 0; x < len(nums); x++ {
		arr[nums[x]+10000]++
	}
	count := 0
	for x := 20000; x >= 0; x-- {
		count = count + arr[x]
		if count >= k {
			return x - 10000
		}

	}
	return -1
}
