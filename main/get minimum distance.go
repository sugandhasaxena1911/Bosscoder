package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 5
	start := 3
	d := getMinDistance(nums, target, start)
	fmt.Println(d)

}
func getMinDistance(nums []int, target int, start int) int {

	if nums[start] == target {
		return 0
	}

	left := start - 1
	right := start + 1

	d1 := 999
	for left >= 0 {
		if nums[left] == target {

			d1 = start - left
			break
		}
		left--
	}
	fmt.Println("dist left ", d1)

	d2 := 999
	for right < len(nums) {
		if nums[right] == target {
			d2 = right - start
			break
		}
		right++
	}
	fmt.Println("dist right ", d2)
	if d1 < d2 {
		return d1

	}
	return d2
}

func getMinDistance2(nums []int, target int, start int) int {
	distance := 999
	if nums[start] == target {
		return 0
	}

	left := start - 1
	right := start + 1

	for left >= 0 || right < len(nums) {
		if left >= 0 {
			if nums[left] == target {

				distance = start - left
				break
			}
		}
		if right < len(nums) {
			if nums[right] == target {

				distance = right - start
				break
			}
		}
		left--
		right++

	}

	return distance

}
