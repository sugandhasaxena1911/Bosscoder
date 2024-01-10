package main

import "fmt"

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	printmatrix(matrix)
	setZeroes(matrix)
	fmt.Println("---------------------------------------------")
	printmatrix(matrix)

	matrix = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	printmatrix(matrix)
	setZeroes(matrix)
	fmt.Println("---------------------------------------------")
	printmatrix(matrix)

}

func setZeroes(matrix [][]int) {

	arraycheck := [][]int{}
	for m := 0; m < len(matrix); m++ {
		for n := 0; n < len(matrix[m]); n++ {

			if matrix[m][n] == 0 {
				arraycheck = append(arraycheck, []int{m, n})
				// making elemnts already parsed =0
				for x := 0; x < m; x++ {
					matrix[x][n] = 0
				}
				for y := 0; y < n; y++ {
					matrix[m][y] = 0
				}
			} else {
				if len(arraycheck) > 0 {
					for z := 0; z < len(arraycheck); z++ {
						if m >= arraycheck[z][0] && n == arraycheck[z][1] {
							matrix[m][n] = 0
						}
						if n >= arraycheck[z][1] && m == arraycheck[z][0] {
							matrix[m][n] = 0
						}

					}
				}
			}

		}
	}
}

func printmatrix(matrix [][]int) {

	for m := 0; m < len(matrix); m++ {
		for n := 0; n < len(matrix[m]); n++ {
			fmt.Print(matrix[m][n])
			fmt.Print(" ")
		}
		fmt.Println("")
	}

}
