package main

import "fmt"

func main() {
	nums := []int{2, 5, 6, 9, 10}
	gcd := findGCD(nums)
	fmt.Println(gcd)

}

func findGCD(nums []int) int {

	// gcd of min & max of array
	mini := nums[0]
	maxi := nums[0]
	for x := 1; x < len(nums); x++ {
		mini = min(mini, nums[x])
		maxi = max(maxi, nums[x])
	}

	return gcd(mini, maxi)

}

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func min(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}

func gcd(n1 int, n2 int) int {
	if n2 == 0 {
		return n1
	}

	return gcd(n2, n1%n2)
}
