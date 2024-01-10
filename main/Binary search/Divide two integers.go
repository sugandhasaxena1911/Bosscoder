package main

import "fmt"

func main() {

	fmt.Println(divide(7, -3))
}

func divide(dividend int, divisor int) int {
	count := 1
	d := divisor
	for divisor < dividend {
		count++
		divisor = d + divisor
		fmt.Println(count, divisor)

	}
	if divisor > dividend {
		count--

	}
	return count
}
