package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 8, 9}
	diff := 2
	//	triplets := arithmeticTriplets(nums, diff)
	//fmt.Println(triplets)
	triplets := arithmeticTriplets2(nums, diff)
	fmt.Println(triplets)

}

//brute force

func arithmeticTriplets(nums []int, diff int) int {
	x, y, z := 0, 0, 0
	total := 0

	for x < len(nums)-2 {
		y = x + 1
		for y < len(nums)-1 {
			fmt.Println("duet ", x, y, nums[y]-nums[x])

			z = y + 1
			if nums[y]-nums[x] == diff {
				//pair found , find triplet
				fmt.Println(" pair found")

				for z < len(nums) {
					fmt.Println("triplet ", x, y, z, nums[z]-nums[y])

					if nums[z]-nums[y] == diff {
						total++
						fmt.Println("total ", total)
						z++
					} else {
						z++
						continue
					}
				}
				y++

			} else {
				y++
			}

		}
		x++
	}

	return total
}

//two pointer approach

func arithmeticTriplets2(nums []int, diff int) int {

	triplets := 0

	i, j, k := 0, 1, 2

	for i < len(nums)-2 && j < len(nums)-1 && k < len(nums) {

		if !findCouple(&i, &j, nums, diff) {
			continue
		}
		// find the other
		fmt.Println("values ", i, j)
		if !findCouple(&j, &k, nums, diff) {
			continue
		}
		i++
		j++
		k++
		triplets++

	}

	return triplets

}

func findCouple(left *int, right *int, nums []int, diff int) bool {
	fmt.Println("inside couple ")
	fmt.Println("inside couple ", *left, *right)
	fmt.Println("inside couple ", nums[*left], nums[*right], diff)

	if diff < nums[*right]-nums[*left] {
		fmt.Println("Inside case 1")
		*left++
		return false
	}
	if diff > nums[*right]-nums[*left] {
		fmt.Println("Inside case 2")

		*right++
		return false
	}

	return true

}
