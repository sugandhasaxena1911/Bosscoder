package main

import "fmt"

func main() {

	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 23
	fmt.Println(searchMatrix(matrix, target))
}

// elemnts on left are smaller & elemnts down are greater than current
func searchMatrix(matrix [][]int, target int) bool {
	row := 0
	col := len(matrix[0]) - 1

	for row < len(matrix) && col >= 0 {
		curr := matrix[row][col]
		if target == curr {
			return true
		}
		if target > curr {
			row++
		} else {
			col--
		}

	}
	return false
}
