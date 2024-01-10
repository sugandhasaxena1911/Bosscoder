package main

import "fmt"

func main() {
	sortColors([]int{1, 0, 0, 1, 2, 1, 0})
	DutchNationalFlag([]int{1, 0, 0, 1, 2, 1, 0})

}

func sortColors(nums []int) {
	l := 0
	r := len(nums) - 1
	for l < r {
		if nums[l] == 0 {
			l++
			continue
		}

		/*if nums[r] == 2 {
			r--
			continue
		}*/
		if nums[l] == 2 {
			nums[r], nums[l] = nums[l], nums[r]
			r--
			continue
		}
		/*if nums[r] == 0 {
			nums[r], nums[l] = nums[l], nums[r]
			l++
			continue
		}*/
		if nums[l] == 1 {
			//nearest zer0
			z := -1
			for x := l + 1; x <= r; x++ {
				if nums[x] == 0 {
					z = x
				}
			}
			if z != -1 {
				//swap
				nums[l], nums[z] = nums[z], nums[l]
				l++
			} else {
				l++
			}
		}

	}

	fmt.Println(nums)

}

func DutchNationalFlag(nums []int) {
	low := 0
	mid := 0
	high := len(nums) - 1

	for mid <= high {
		if nums[mid] == 0 {
			nums[mid], nums[low] = nums[low], nums[mid]
			mid++
			low++
		} else if nums[mid] == 1 {
			mid++
		} else if nums[mid] == 2 {
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}

	}
	fmt.Println(nums)
}
