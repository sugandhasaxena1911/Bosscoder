package main

import "fmt"

func main() {
	columnNumber := 703
	st := convertToTitleSimple(columnNumber)
	fmt.Println(st)

}

var arrref []string

func convertToTitleSimple(columnNumber int) string {
	st := ""
	v := 0
	arr := []int{}

	for columnNumber > 0 {
		v = (columnNumber - 1) % 26
		columnNumber = (columnNumber - 1) / 26
		arr = append(arr, v)
		fmt.Println("v ", v)
		fmt.Println("columnnumber ", columnNumber)
	}

	fmt.Println(arr)
	for x := len(arr) - 1; x >= 0; x-- {
		// Anotherway
		st = st + string(65+arr[x])
	}
	return st

}

func convertToTitle(columnNumber int) string {
	st := ""
	v := 0
	arrref = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	arr := []int{}

	for columnNumber > 0 {
		if columnNumber%26 == 0 {
			v = (columnNumber - 26) % 26
			columnNumber = (columnNumber - 26) / 26

		} else {
			v = columnNumber % 26
			columnNumber = columnNumber / 26
		}
		fmt.Println("v ", v)
		arr = append(arr, v)
		fmt.Println("columnnumber ", columnNumber)
	}

	fmt.Println(arr)
	for x := len(arr) - 1; x >= 0; x-- {
		//st = st + makeString(arr[x])
		// Anotherway
		if arr[x] == 0 {
			st = st + string(arr[x]+122)

		} else {
			st = st + string(arr[x]+64)
		}
	}
	return st

}

func makeString(n int) string {
	if n == 0 {
		return arrref[25]
	} else {
		return arrref[n-1]
	}
}

func convertToTitleV2(columnNumber int) string {
	st := ""
	v := 0
	arr := []int{}

	for columnNumber > 0 {
		if columnNumber%26 == 0 {
			v = (columnNumber - 26) % 26
			columnNumber = (columnNumber - 26) / 26

		} else {
			v = columnNumber % 26
			columnNumber = columnNumber / 26
		}
		fmt.Println("v ", v)
		arr = append(arr, v)
		fmt.Println("columnnumber ", columnNumber)
	}

	fmt.Println(arr)
	for x := len(arr) - 1; x >= 0; x-- {
		// Anotherway
		if arr[x] == 0 {
			st = st + string(arr[x]+122)

		} else {
			st = st + string(arr[x]+64)
		}
	}
	return st

}
