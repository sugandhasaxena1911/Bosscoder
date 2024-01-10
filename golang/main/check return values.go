package main

import "fmt"

func main() {
	s := test(5)
	fmt.Println("Yo ", s)

}

func test(n int) (str string) {
	str = "Hey"
	fmt.Println("n ", n, str)
	return
}
