package main

import "fmt"

func main() {

	n := countPrimes(50)
	fmt.Println(n)
	n = CountPrimeSieveMethod(50)
	fmt.Println(n)

}

func countPrimes(n int) int {
	cnt := 0
	for x := 2; x < n; x++ {
		if IsPrimeOptimized(x) {
			cnt++
		}
	}

	return cnt

}

func IsPrimeOptimized(n int) bool {
	if n == 1 {
		return false
	}

	for x := 2; x*x <= n; x++ {
		if n%x == 0 {
			return false
		}
	}
	return true
}

func CountPrimeSieveMethod(n int) int {
	prime := 0
	// all entries in p are marked 0
	p := make([]int, n+1)

	for x := 2; x*x <= n; x++ {
		if p[x] == 0 {
			for y := x * x; y <= n; y = y + x {
				p[y] = 1
			}

		}
	}

	for x := 2; x <= n; x++ {
		if p[x] == 0 {
			prime++
		}
	}
	return prime
}
