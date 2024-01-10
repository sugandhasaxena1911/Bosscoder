package main

import "fmt"

func main() {
	mat := [][]int{{1, 2}, {3, 4}}
	r := 4
	c := 1
	newmat := matrixReshape(mat, r, c)
	fmt.Println(newmat)

}
func matrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	n := len(mat[0])
	if m*n != r*c {
		return mat
	}

	newmat := make([][]int, r)
	for x := 0; x < r; x++ {
		newmat[x] = make([]int, c)
	}

	// start reshaping
	row := 0
	col := 0
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			newmat[row][col] = mat[x][y]
			col++
			if col > c-1 {
				col = 0
				row++
			}

		}
	}
	return newmat
}
