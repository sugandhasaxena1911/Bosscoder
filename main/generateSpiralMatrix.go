package main

import "fmt"

func main() {
	matrix := generateMatrix(5)
	printmatrix(matrix)
}

func generateMatrix(n int) [][]int {
	mat := make([][]int, n)
	fmt.Println(mat)
	for x := 0; x < n; x++ {
		mat[x] = make([]int, n)

	}
	num := 1
	l := 0
	r := n - 1
	t := 0
	b := n - 1

	for num <= n*n {

		for x := l; x <= r; x++ {
			mat[l][x] = num
			fmt.Println(l, x, num)
			num++
		}
		r--

		for x := t + 1; x <= b; x++ {
			mat[x][b] = num
			fmt.Println(x, b, num)

			num++
		}
		b--

		for x := r; x >= l; x-- {
			mat[r+1][x] = num
			fmt.Println(r, x, num)

			num++
		}
		l++

		for x := b; x >= t+1; x-- {
			mat[x][t] = num
			fmt.Println(x, t, num)
			num++
		}
		t++
	}

	return mat
}

func printmatrix(matrix [][]int) {

	for m := 0; m < len(matrix); m++ {
		for n := 0; n < len(matrix[m]); n++ {
			if matrix[m][n] < 10 {
				fmt.Print(matrix[m][n], "  ")
			} else {
				fmt.Print(matrix[m][n], " ")

			}
			fmt.Print(" ")
		}
		fmt.Println("")
	}

}
