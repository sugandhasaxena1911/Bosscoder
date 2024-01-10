package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 8}
	queries := [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}
	result := xorQueries(arr, queries)
	fmt.Println(result)

}

// brute force
func xorQueries(arr []int, queries [][]int) []int {
	result := []int{}
	for x := 0; x < len(queries); x++ {
		r := arr[queries[x][0]]

		for y := queries[x][0] + 1; y <= queries[x][1]; y++ {

			r = r ^ arr[y]
		}
		result = append(result, r)

	}
	return result
}
