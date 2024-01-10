package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(CheckOddEvenFactors(25))
	fmt.Println(CheckOddEvenFactors(39))

}

// O(1)
func CheckOddEvenFactors(n int) string {

	sqroot := math.Sqrt(float64(n))
	a := int(sqroot)
	if a*a == n {
		return "Odd"
	} else {
		return "Even"
	}
}
