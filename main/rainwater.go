package main

import "fmt"

func main() {

	//water := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	water := trap([]int{4, 2, 0, 3, 2, 5})

	fmt.Println("the total water is ", water)
}

func trap(height []int) int {

	left := calLeftMax(height)
	right := calRightMax(height)
	water := 0
	level := 0

	for i := range height {
		level = max(0, (min(left[i], right[i]) - height[i]))
		fmt.Println("level at index  ", i, " ", level)
		water = water + level
	}

	return water

}

func calLeftMax(height []int) []int {
	maxval := 0
	left := make([]int, len(height))
	//fmt.Println("left ", left)

	for i := 1; i < len(height); i++ {
		fmt.Println(maxval, height[i-1])
		maxval = max(maxval, height[i-1])
		left[i] = maxval
	}
	fmt.Println("left max array ", left)
	return left
}

func calRightMax(height []int) []int {
	maxval := 0
	right := make([]int, len(height))
	//fmt.Println("right ", right)
	for i := len(height) - 2; i >= 0; i-- {
		maxval = max(maxval, height[i+1])
		right[i] = maxval
	}
	fmt.Println("rightmax array ", right)

	return right

}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b

}
func min(a int, b int) int {
	if a < b {
		return a
	}

	return b

}
