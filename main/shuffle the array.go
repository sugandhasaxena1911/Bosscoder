package main

import "fmt"

func main() {
	arr := []int{2, 5, 1, 3, 4, 7}
	result := shuffle(arr, 3)
	fmt.Println(result)

}
func shuffle(nums []int, n int) []int {
	//using second array
	arr2 := make([]int, n*2)
	y := n
	for x := 0; x < (n*2)-1; x = x + 2 {
		arr2[x] = nums[x/2]
		arr2[x+1] = nums[y]
		y++
	}

	return arr2
}

func shuffle2(nums []int, n int) []int {
	//using same array : rearranging
	arr2 := make([]int, n*2)
	y := n
	for x := 0; x < (n*2)-1; x = x + 1 {
		if x%2 ==1{
			nums[x],nums[y]= nums[y],nums[x]
			y++
		}else{
			nums[x] = nums[y--]
		}
	}

	return arr2
}

/*
0  1  2  3	4  5  6   7
x1 x2 x3 x4 y1 y2 y3 y4
x1 y1 x2 y2 x3 y3 x4 y4
x1 y1       x2 
         y2    x4
   1      3    5     7
*/
