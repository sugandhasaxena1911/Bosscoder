package main

import "fmt"

func main() {
	factors, count := FactorsOfN(45)
	fmt.Println(factors, count)

}

// O(sqrt(n))
func FactorsOfN(n int) ([]int, int) {
	arr := []int{}
	for x := 1; x*x <= n; x++ {
		fmt.Println(x)
		if n%x == 0 {

			if n/x == x {
				arr = append(arr, x)
			} else {
				arr = append(arr, x, n/x)
			}
		}

	}
	return arr, len(arr)
}
