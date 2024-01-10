package main

import "fmt"

////CHECK FOR OPTIMAL SOLN
func main() {
	arr := []int{1, 4, 2, 5, 3}
	res := sumOddLengthSubarrays2(arr)
	fmt.Println(res)

	arr = []int{10, 11, 12}
	res = sumOddLengthSubarrays2(arr)
	fmt.Println(res)

}

func sumOddLengthSubarrays(arr []int) int {
	sum := 0
	arr2 := make([]int, len(arr))

	//create prefix sum
	for i, v := range arr {
		sum = sum + v
		arr2[i] = sum
	}
	fmt.Println(arr2)
	sum = 0
	oddidx := 1
	for oddidx <= len(arr2) {
		fmt.Println("oddidx ", oddidx)
		for i := oddidx - 1; i < len(arr2); i++ {
			if i-oddidx < 0 {
				sum = sum + (arr2[i])
			} else {
				sum = sum + (arr2[i] - arr2[i-oddidx])
			}
			fmt.Println("sum ", sum)
		}
		oddidx = oddidx + 2
	}
	return sum
}

func sumOddLengthSubarrays2(arr []int) int {
	sum := 0

	sum = 0
	oddidx := 1
	for oddidx <= len(arr) {
		for x := 0; x+oddidx <= len(arr); x++ {

			sub := arr[x : x+oddidx]
			fmt.Println(sub)
			//traverse sub array
			for y := 0; y < len(sub); y++ {
				sum = sum + sub[y]
			}

		}

		oddidx = oddidx + 2
	}
	return sum
}
