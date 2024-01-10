package main

import (
	"fmt"
	"strings"
)

// gives time limit exceeded in leetcode
// Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
func main() {

	s1 := "ab"
	s2 := "eidbaooo"
	b := checkInclusion(s1, s2)
	fmt.Println(b)

}
func checkInclusion(s1 string, s2 string) bool {
	return PermutationsOfString(s1, "", s2)

}

func PermutationsOfString(s string, answer string, s2 string) bool {
	if len(s) == 0 {
		b := strings.Contains(s2, answer)
		fmt.Println("is substring ", answer, " present ?? ", b)
		return b

	}
	found := false
	// we make two strings question string & answer string
	for i, v := range s {
		//fmt.Println("Character ", string(v))
		//answer = answer + string(v)
		newQ := s[0:i] + s[i+1:]
		//fmt.Println("newQ n answer ", newQ, answer)

		found = PermutationsOfString(newQ, answer+string(v), s2)
		if found {
			break
		}
	}

	return found
}
