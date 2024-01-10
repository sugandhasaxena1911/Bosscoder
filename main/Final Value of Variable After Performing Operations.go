package main

import "fmt"

func main() {
	arr := []string{"--X", "X++", "X++"}
	val := finalValueAfterOperations(arr)
	fmt.Println(val)
	val = finalValueAfterOperations2(arr)
	fmt.Println(val)
}

func finalValueAfterOperations(operations []string) int {
	value := 0
	for _, v := range operations {
		if v == "++X" || v == "X++" {
			value++
		} else {
			value--
		}
	}
	return value
}

func finalValueAfterOperations2(operations []string) int {
	value := 0
	for _, v := range operations {
		switch v {
		case "++X", "X++":
			value++
		default:
			value--

		}
	}
	return value
}
