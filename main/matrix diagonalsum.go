package main

import "fmt"

func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	sum := diagonalSum(mat)
	fmt.Println(sum)

}

func diagonalSum(mat [][]int) int {
	sum := 0
	c := len(mat) - 1
	for x := 0; x < len(mat); x++ {
		sum = sum + mat[x][x]
		if x != c {
			sum = sum + mat[x][c]
		}
		c--
	}

	return sum
}
