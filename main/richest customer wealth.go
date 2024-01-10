package main

import "fmt"

func main() {
	accounts := [][]int{{1, 2, 3}, {3, 2, 1}}

	wealth := maximumWealth(accounts)
	fmt.Println(wealth)

}

func maximumWealth(accounts [][]int) int {
	richest := 0
	for x := 0; x < len(accounts); x++ {
		wealth := 0
		for y := 0; y < len(accounts[x]); y++ {

			wealth = wealth + accounts[x][y]

		}
		richest = max(richest, wealth)
	}
	return richest
}

func max(n1 int, n2 int) int {

	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}
