package main

import "fmt"

func main() {
	PermutationsOfString("abc", "")

}

func PermutationsOfString(s string, answer string) {
	if len(s) == 0 {
		fmt.Println("output", answer)

		return
	}
	// we make two strings question string & answer string
	for i, v := range s {
		//fmt.Println("Character ", string(v))
		//answer = answer + string(v)
		newQ := s[0:i] + s[i+1:]
		//fmt.Println("newQ n answer ", newQ, answer)

		PermutationsOfString(newQ, answer+string(v))
	}
}
