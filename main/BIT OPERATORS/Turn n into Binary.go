package main

import "fmt"

func main() {

	binary := ToBinary(9)
	fmt.Println(binary)

}

func ToBinary(n int) []int {

	binary := []int{}

	for n > 0 {
		binary = append(binary, n%2)
		n = n / 2
	}

	ReverseArray(binary)
	return binary
}

func ReverseArray(n []int) {
	l := 0
	r := len(n) - 1
	for l < r {
		n[l], n[r] = n[r], n[l]
		l++
		r--
	}

}
