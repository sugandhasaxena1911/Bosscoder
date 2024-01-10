package main

import "fmt"

func main() {
	hours := []int{98}
	target := 5
	fmt.Println(numberOfEmployeesWhoMetTarget(hours, target))

}
func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	n := 0
	for _, v := range hours {
		if v >= target {
			n++
		}
	}
	return n
}
