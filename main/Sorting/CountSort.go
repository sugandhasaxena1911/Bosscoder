package main

import "fmt"

func main() {
	nums := []int{7, 3, 5, 7}
	n := CountSort(nums)
	fmt.Println(n)
}

func CountSort(nums []int) []int {
	fmt.Println(nums)

	max := FindMax(nums)
	fmt.Println(max)
	arr := make([]int, max+1)
	fmt.Println("array sum ", arr)

	numssort := make([]int, len(nums))

	for x := 0; x < len(nums); x++ {
		arr[nums[x]]++
	}
	fmt.Println("array sum ", arr)

	// calculate prefix sum
	for x := 1; x < len(arr); x++ {
		arr[x] = arr[x] + arr[x-1]
	}
	fmt.Println("array PREFIX sum ", arr)

	for x := 0; x < len(nums); x++ {
		//nums[x]=key,   index of key  in sorted array =  arr[key]-1 & then nums[x]--
		key := nums[x]
		index := arr[key] - 1
		numssort[index] = key
		arr[key]--
	}

	return numssort
}

func FindMax(nums []int) int {

	max := nums[0]
	for x := 1; x < len(nums); x++ {
		if nums[x] > max {
			max = nums[x]
		}

	}

	return max
}
