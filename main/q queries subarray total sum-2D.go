package main

import "fmt"

func main() {
	nums := [][]int{{1, 2, 3, 4, 6},
		{5, 3, 8, 1, 2},
		{4, 6, 7, 5, 5},
		{2, 4, 8, 9, 4}}

	//elements of qnums define the top left & bottom right index of the submatrix
	// 0,1 : index of top left 2,3: index of right bottom
	qnums := [][]int{{0, 0, 1, 1}, {2, 2, 3, 4}, {1, 2, 3, 3}, {0, 1, 1, 2}}
	sum := RangeSumQueriesV2(nums, qnums)
	fmt.Println("Total of all submatrices ", sum)

}

//BRUTE FORCE
// Complexity : O(q*m*n)  if q=m*n then O(n*4)
func RangeSumQueries(nums [][]int, qnums [][]int) int {
	q := len(qnums)
	sum := 0

	for x := 0; x < q; x++ {
		tlX := qnums[x][0]
		tlY := qnums[x][1]
		brX := qnums[x][2]
		brY := qnums[x][3]
		fmt.Println("indexes (", tlX, ",", tlY, ")")
		fmt.Println("indexes (", brX, ",", brY, ")")

		// subarray left index & right index calculated
		//parse over 2Darray
		s := 0
		for x := tlX; x <= brX; x++ {
			for y := tlY; y <= brY; y++ {
				fmt.Println("element ", nums[x][y])
				s = s + nums[x][y]

			}
		}
		fmt.Println("submatrix sum ", s)
		sum = sum + s
	}
	return sum

}

// Using PRE-PROCESSING Time O(q)+O(n)=O(n+q)  if n==q   O(2n)=O(n), extra space : O(n)
func RangeSumQueriesV2(nums [][]int, qnums [][]int) int {
	q := len(qnums)
	fmt.Println("IGNORE ", q)
	sum := 0

	// Create presum matrix
	presum := make([][]int, len(nums))
	for x := 0; x < len(nums); x++ {
		presum[x] = make([]int, len(nums[x]))
	}
	//Set first row
	for x := 1; x < len(nums[0]); x++ {
		nums[0][x] = nums[0][x] + nums[0][x-1]
		fmt.Println("VAL ", nums[0][x])
	}
	//Set first column
	for x := 1; x < len(nums); x++ {
		nums[x][0] = nums[x][0] + nums[x-1][0]
		fmt.Println("VAL ", nums[x][0])
	}
	//Set other values
	for x := 1; x < len(nums); x++ {
		for y := 1; y < len(nums[0]); y++ {
			nums[x][y] = nums[x][y] + nums[x][y-1] + nums[x-1][y] - nums[x-1][y-1]
			fmt.Println("(", x, ",", y, ")", "  ", nums[x][y])

		}
	}

	// Start calculating sum of submatrices
	for x := 0; x < q; x++ {
		tlX := qnums[x][0]
		tlY := qnums[x][1]
		brX := qnums[x][2]
		brY := qnums[x][3]
		fmt.Println("TOP (", tlX, ",", tlY, ")")
		fmt.Println("BOTTOM (", brX, ",", brY, ")")

		// subarray left index & right index calculated
		//parse over 2Darray
		s := nums[brX][brY]
		if tlX > 0 {
			s = s - nums[tlX-1][brY]
		}
		if tlY > 0 {
			s = s - nums[brX][tlY-1]
		}
		if tlX > 0 && tlY > 0 {
			s = s + nums[tlX-1][tlY-1]
		}
		fmt.Println("s ", s)
		sum = sum + s
	}

	return sum
}
