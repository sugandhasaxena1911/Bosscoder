package main

import "fmt"

func main() {
	n := commonFactors(25, 30)
	fmt.Println(n)
}

func commonFactors(a int, b int) int {
	n := 0
	for x := 1; x <= min(a, b); x++ {
		fmt.Println(x)
		if a%x == 0 && b%x == 0 {
			n++
		}
	}
	return n
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
