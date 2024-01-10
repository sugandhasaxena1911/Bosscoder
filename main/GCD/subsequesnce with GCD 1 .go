package main

import "fmt"

func main() {

	nums := []int{15, 42, 8}
	IsGCD1 := IsSubArraywithGCD1(nums)
	fmt.Println(IsGCD1)

}

//if subrray gcd==1 then whole array gcd ==1 .
func IsSubArraywithGCD1(nums []int) bool {

	result := nums[0]
	for x := 1; x < len(nums); x++ {
		result = gcd(result, nums[x])
		if result == 1 {
			break
		}
	}
	if result == 1 {
		return true
	} else {
		return false
	}
}

func gcd(n1 int, n2 int) int {
	if n2 == 0 {
		return n1
	}

	return gcd(n2, n1%n2)
}
