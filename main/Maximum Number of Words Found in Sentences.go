package main

import (
	"fmt"
	"strings"
)

func main() {

	sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	count := mostWordsFound(sentences)
	fmt.Println(count)

}
func mostWordsFound(sentences []string) int {
	words := []string{}
	maximum := 0
	for _, v := range sentences {
		words = strings.Split(v, " ")
		maximum = max(maximum, len(words))

	}
	return maximum
}

func max(n1 int, n2 int) int {

	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}
