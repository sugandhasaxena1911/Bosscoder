package main

import "fmt"

func main() {

	// variable  : elements has 0 value
	fmt.Println("STATIC ARRAY WITH ZERO VALUE ELEMENTS -USING VAR")
	var arr [5]int
	fmt.Println(arr, len(arr), cap(arr))
	//assign value
	arr[1] = 5
	fmt.Println(arr)

	// using dyanmic , EMPLTY SLICE
	fmt.Println("DYNAMIC ARRAY , NIL SLICE-USING VAR")
	var arr2 []int
	fmt.Println(arr2, len(arr2), cap(arr2))
	//assign value
	//arr2[1] = 5 // INDEX OUT OF RANGE
	arr2 = append(arr2, 5, 6, 7)
	fmt.Println(arr2, len(arr2), cap(arr2))

	// Composite literal
	var arr3 []int
	arr3 = []int{1, 3, 4, 4}
	fmt.Println("composite ", arr3, len(arr3), cap(arr3))

	//slicing   // RETURNS NEW ARRAY
	fmt.Println("Slicing returns new array ", arr2[1:3], arr2)

	fmt.Println("DYNAMIC ARRAY , USING MAKE with LENGTH")
	arr4 := make([]int, 5)
	fmt.Println(arr4, len(arr4), cap(arr4))

	fmt.Println("DYNAMIC ARRAY , USING MAKE with LENGTH & CAPACITY")
	arr5 := make([]int, 5, 7)
	fmt.Println(arr5, len(arr5), cap(arr5))

	fmt.Println("ADD ELEMENTS BEYOND CAPACITY")
	arr5 = append(arr5, 9, 6, 1) // CAPACITY DOUBLES
	fmt.Println("AFTER ", arr5, len(arr5), cap(arr5))

	fmt.Println("DYNAMIC ARRAY , WITHOUT MAKE & var")
	arr6 := []int{1, 2, 3, 4}
	fmt.Println(arr6, len(arr6), cap(arr6))

	arr7empty := []int{}
	fmt.Println(arr7empty, len(arr7empty), cap(arr7empty))

	var m int
	fmt.Println(m)

}
