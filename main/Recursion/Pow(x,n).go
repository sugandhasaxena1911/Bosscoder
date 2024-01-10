package main

import "fmt"

func main() {
	x := 0.1
	n := 2147483647
	//ans := myPow(x, n)
	//fmt.Println("Answer is ", ans)

	ans2 := myPowV2(x, n)
	fmt.Println("Answer is ", ans2)
}

// will give memory issue for higher values
func myPow(x float64, n int) float64 {
	flag := 1
	if n < 0 {
		fmt.Println("n is ", n)
		flag = -1
		n = n * flag
	}

	if n == 0 {
		fmt.Println("n is 0")
		return 1
	}
	if flag == 1 {
		fmt.Println("n is >0", n)
		return x * myPow(x, n-1)
	} else {
		fmt.Println("n is <0", n)
		b := x * myPow(x, n-1)
		return 1 / b
	}

}

func myPowV2(x float64, n int) float64 {
	flag := 1
	if n < 0 {
		fmt.Println("n is ", n)
		flag = -1
		n = n * flag
	}

	if n == 0 {
		fmt.Println("n is 0, return 1")
		return 1
	}
	if flag == 1 {
		fmt.Println("n is >0", n)
		hp := myPowV2(x, n/2)
		var b float64

		if n%2 == 0 {
			b = hp * hp
			fmt.Println("b ", b)
			return b
		} else {
			b = hp * hp * x
			fmt.Println("b ", b)
			return b
		}
	} else {
		fmt.Println("n is <0", n)
		hp := myPowV2(x, n/2)
		var b float64
		if n%2 == 0 {
			b = hp * hp
		} else {
			b = hp * hp * x
		}
		return 1 / b
	}

}
