package main

import "fmt"

func main() {
	left := 40
	right := 57

	nums := selfDividingNumbers(left, right)
	fmt.Println(nums)
}
func selfDividingNumbers(left int, right int) []int {
	nums := []int{}
	for x := left; x <= right; x++ {
		fmt.Println(x)
		n := x
		for n > 0 {
			ld := n % 10
			if ld == 0 {
				break
			}
			if x%ld != 0 {
				break
			}
			n = n / 10
		}
		if n == 0 {
			nums = append(nums, x)
		}
		if x%10 == 9 {
			x++
		}
	}
	return nums
}
