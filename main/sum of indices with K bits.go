package main

import "fmt"

func main() {
	num := []int{4, 3, 2, 1}
	k := 2
	sum := sumIndicesWithKSetBits(num, k)
	fmt.Println(sum)

}

func sumIndicesWithKSetBits(nums []int, k int) int {
	sum := 0
	total := 0
	for i, v := range nums {
		total = TotalKbitsBinary(i)
		fmt.Println("total is ", total)
		if total == k {
			sum = sum + v
		}
	}
	return sum
}

func TotalKbitsBinary(n int) int {
	fmt.Println("number is ", n)
	bit := 0
	total := 0
	for n > 0 {
		bit = n % 2
		fmt.Println(bit)
		n = n / 2
		if bit == 1 {
			total++
		}
	}
	return total
}
