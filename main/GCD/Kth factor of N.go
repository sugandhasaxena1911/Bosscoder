package main

import "fmt"

func main() {
	n := 4
	k := 4
	//num := kthFactor(n, k)
	//fmt.Println(num)

	num := kthFactorOptimized(n, k)
	fmt.Println(num)

}

//brute
func kthFactor(n int, k int) int {
	if k == 1 {
		return 1
	}
	nth := 1
	for x := 2; x <= n; x++ {
		if n%x == 0 {
			nth++
		}
		if nth == k {
			return x
		}
	}
	return -1
}

// Otimized
func kthFactorOptimized(n int, k int) int {
	factors := []int{}
	count := 0
	for x := 1; x*x <= n; x++ {
		if n%x == 0 {
			if x == n/x {
				count++
			} else {
				count = count + 2
			}
			factors = append(factors, x)
		}
		fmt.Println(len(factors), x)
		if k == len(factors) {
			return x
		}
	}

	l := len(factors)
	fmt.Println("factors array length ", l)
	fmt.Println("totsl fsctors  ", count)

	if k <= count {
		if count%2 == 0 { //not a perfct square
			return n / factors[l-(k-l)] // k-l=difference
		} else {
			return n / factors[l-(k-l)-1]
		}
	}
	return -1
}
