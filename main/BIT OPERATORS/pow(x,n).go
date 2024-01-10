package main

import "fmt"

func main() {
	res := myPow(2.1, 3)
	fmt.Println(res)

}
func myPow(x float64, n int) float64 {
	pow := float64(1)
	a := x
	minus := false
	if n < 0 {
		minus = true
		n = -n
	}
	// the below commented code  works, copied from leetcode
	for n > 0 {
		if (n & 1) != 0 { // ie if n is odd ,
			pow = pow * a
		}
		a = a * a // each time square since we are divind exponent by 2
		n >>= 1   // divide n by 2
	}

	if minus {
		pow = 1 / pow
	}
	fmt.Println("pow is ", pow)

	return pow
}
