package main

import (
	"fmt"
	"math"
)

func main() {

	nums := []int{13, 25, 83, 77}
	arr := separateDigits(nums)
	fmt.Println(arr)
	arr2 := separateDigits(nums)
	fmt.Println(arr2)

}

func separateDigits(nums []int) []int {
	arr := []int{}
	//345/100 = 3
	for _, v := range nums {
		n := []int{}

		for v > 0 {
			lastdg := v % 10
			n = append(n, lastdg)
			v = v / 10

		}

		for i := len(n) - 1; i >= 0; i-- {
			arr = append(arr, n[i])
		}

	}
	return arr
}

// in the next appoach , if we divide 3456/1000 = 3 we get first digit  .
// So , then we can easily keep on appending separte  digits in new array
func separateDigits2(nums []int) []int {
	arr := []int{}
	//345/100 = 3
	for _, v := range nums {
		lg := math.Log10(float64(v))
		for v > 0 {
			pw := int(math.Pow10(int(lg)))
			firstdg := v / pw
			arr = append(arr, firstdg)
			v = v % pw
		}
	}
	return arr
}
