package main

import "fmt"

func main() {
	nums := []int{4, 5, 5, 4, 4, 3, 6, 3, 5, 3}
	n := singleNumber(nums)
	fmt.Println(n)

}

func singleNumber(nums []int) int {
	n := 0
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	for key, _ := range m {

		if m[key] == 1 {
			n = key
			break
		}
	}
	return n
}
