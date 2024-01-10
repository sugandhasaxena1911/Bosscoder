package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 2}
	n := sumOfUnique(nums)
	fmt.Println(n)

	n = sumOfUnique2(nums)
	fmt.Println("sum is ", n)
}

func sumOfUnique(nums []int) int {
	sum := 0

	m1 := make(map[int]int)
	for _, v := range nums {
		fmt.Println(v, m1[v])

		if _, ok := m1[v]; !ok {
			fmt.Println("here")
			sum = sum + v

		} else {
			if m1[v] == 1 {
				sum = sum - v

			}

		}
		m1[v]++

	}

	return sum

}

func sumOfUnique2(nums []int) int {
	sort.Ints(nums)
	//12223 4
	sum := 0
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	if nums[0] != nums[1] {
		sum = sum + nums[0]
	}

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
			sum = sum + nums[i]
		}
	}
	if nums[l-1] != nums[l-2] {
		sum = sum + nums[l-1]
	}
	return sum

}
