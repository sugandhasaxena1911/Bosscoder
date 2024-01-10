package main

import "fmt"

func main() {
	n := 5

	arr := sumZero(n)
	fmt.Println(arr)

}
func sumZero(n int) []int {
	arr := make([]int, n)
	l := n / 2
	for i := 0; i < l; i++ {
		arr[i] = i + 1
		arr[i+l] = -(i + 1)
	}

	return arr

}
