package main

import "fmt"

func main() {
	original := []int{1, 2, 3, 4}
	m := 2
	n := 2
	matrix := construct2DArray(original, m, n)
	fmt.Println(matrix)

	matrix = construct2DArrayV2(original, m, n)
	fmt.Println(matrix)

}

func construct2DArray(original []int, m int, n int) [][]int {
	matrix := make([][]int, m)
	if len(original) != n*m {
		return [][]int{}
	}

	row := -1
	for x := 0; x < m; x++ {
		matrix[x] = make([]int, n)
	}
	for x := 0; x < len(original); x++ {
		if x%n == 0 {
			row++
		}
		matrix[row][x%n] = original[x]
	}
	return matrix
}

func construct2DArrayV2(original []int, m int, n int) [][]int {
	matrix := [][]int{}
	if len(original) != n*m {
		return matrix
	}

	for x := 0; x < m; x++ {

		matrix = append(matrix, original[0:n])
		original = original[n:]
	}
	return matrix
}
