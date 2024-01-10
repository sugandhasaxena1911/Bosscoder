package main

import "fmt"

func main() {
	prime := PrimeNumbers(50)
	fmt.Println(prime)

}

func PrimeNumbers(n int) []int {
	prime := []int{}
	for x := 2; x <= n; x++ {
		if IsPrimeOptimized(x) {
			prime = append(prime, x)
		}

	}
	return prime

}

// Approach 1 : BRUTE

func IsPrime(n int) bool {

	for x := 2; x <= n-1; x++ {
		if n%x == 0 {
			return false
		}
	}
	return true
}

//O(sqrt(n))
//Approach 2 : factors are present in pairs .
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
