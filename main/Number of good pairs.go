package main

import (
	"fmt"
	"sort"
)

//A pair (i, j) is called good if nums[i] == nums[j] and i < j.

func main() {
	arr := []int{1, 2, 3, 1, 1, 3, 1, 3, 2}
	result := numIdenticalPairs(arr)
	fmt.Println(result)
	fmt.Println(numIdenticalPairs2(arr))
	fmt.Println(numIdenticalPairs3(arr))
	fmt.Println("new")
	arr = []int{1, 2, 3, 1, 1, 3, 1, 3, 2}

	fmt.Println(numIdenticalPairs4(arr))

}
func numIdenticalPairs(nums []int) int {
	//brute force
	total := 0
	for i, _ := range nums {
		for y := i + 1; y < len(nums); y++ {
			if nums[i] == nums[y] {
				total++
			}
		}
	}

	return total

}
func numIdenticalPairs2(nums []int) int {
	//sort
	cnt := 0
	total := 0
	sort.Ints(nums)
	fmt.Println("nums ", nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] == nums[i] {
			cnt++
		} else {
			total = total + (cnt*(cnt+1))/2
			fmt.Println("total ", total)
			cnt = 0
		}
	}
	total = total + (cnt*(cnt+1))/2

	return total

}

func numIdenticalPairs3(nums []int) int {
	//map
	total := 0
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	fmt.Println("map ", m)
	for _, val := range m {
		total = total + ((val-1)*val)/2

	}
	return total
}

func numIdenticalPairs4(nums []int) int {
	//map
	total := 0
	//	arr := []int{1, 2, 3, 1, 1, 3, 1, 3, 2}

	m := map[int]int{}
	for _, v := range nums {
		total = total + m[v]
		fmt.Println(v, m[v], total)
		m[v]++
	}
	return total

}
