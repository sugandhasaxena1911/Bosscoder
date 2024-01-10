package main

import "fmt"

func main() {
	mat := [][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}
	cnt := numSpecial(mat)
	fmt.Println(cnt)

}

func numSpecial(mat [][]int) int {
	count := 0
	m := len(mat)
	n := len(mat[0])
	row := make([]int, m)
	col := make([]int, n)

	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if mat[x][y] == 1 {
				row[x]++
				col[y]++
			}
		}

	}
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if mat[x][y] == 1 && row[x] == 1 && col[y] == 1 {
				count++
			}
		}

	}

	return count
}
