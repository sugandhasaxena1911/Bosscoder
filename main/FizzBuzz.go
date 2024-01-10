package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(5))

}
func fizzBuzz(n int) []string {
	arr := make([]string, n)
	k := 0
	for x := 0; x < n; x++ {
		k = x + 1
		if k%3 == 0 && k%5 == 0 {
			arr[x] = "FizzBuzz"
		} else if k%3 == 0 {
			arr[x] = "Fizz"
		} else if k%5 == 0 {
			arr[x] = "Buzz"

		} else {

			arr[x] = strconv.Itoa(k)

		}
	}
	return arr
}
