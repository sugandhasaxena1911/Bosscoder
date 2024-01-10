package main

import "fmt"

func main() {

	nums := []int{15, 45, 30}
	gcd := gcdofarray(nums)
	fmt.Println(gcd)

}

func gcdofarray(nums []int) int {

	result := nums[0]
	for x := 1; x < len(nums); x++ {
		result = gcd(result, nums[x])
		if result == 1 {
			return result
		}
	}
	return result
}

func gcd(n1 int, n2 int) int {
	if n2 == 0 {
		return n1
	}

	return gcd(n2, n1%n2)
}
