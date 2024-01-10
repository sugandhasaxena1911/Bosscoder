package main

import "fmt"

func main() {
	prime := SieveOfEran(50)
	fmt.Println(prime)

}

// Create a boolean array "prime[0..n]" and
// initialize all entries it as true.
// start from entries , if entry is true then mark false for for entries >=entry*entry & multiple of entry
// at the end all entries as true are prime (ie untouched entries)
//Time complexity ; sqrt(n)
func SieveOfEran(n int) []int {
	prime := []int{}
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
			prime = append(prime, x)
		}
	}
	return prime
}
