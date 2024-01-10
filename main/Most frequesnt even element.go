package main

import "fmt"

func main() {
	nums := []int{8154, 9139, 8194, 3346, 5450, 9190, 133, 8239, 4606, 8671, 8412, 6290}
	n := mostFrequentEven(nums)
	fmt.Println(n)
	m := mostFrequentEven(nums)
	fmt.Println(m)

}

func mostFrequentEven(nums []int) int {
	m1 := make(map[int]int)
	freq, n := 0, -1

	for _, v := range nums {
		if v%2 == 0 {
			m1[v]++
		}

	}
	for k, v := range m1 {
		fmt.Println(k, v)
		if v > freq {
			freq = v
			n = k
		}

	}

	for k, v := range m1 {
		fmt.Println(k, v)
		if v == freq && k < n {
			n = k
		}

	}

	return n
}

func mostFrequentEven2(nums []int) int {
	m1 := make(map[int]int)
	freq, result := 0, -1

	for _, v := range nums {
		if v%2 == 0 {
			m1[v]++
		}

	}
	for k, v := range m1 {
		fmt.Println(k, v)
		if v >= freq {
			if v == freq {
				result = min(result, k)

			} else {
				result = k
			}
			freq = v
		}

	}

	return result
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}
