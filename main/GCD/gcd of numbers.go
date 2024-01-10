package main

import "fmt"

func main() {
	n1 := 18
	n2 := 27

	gcd := gcdofnums(n1, n2)
	fmt.Println(gcd)

	gcd = gcdEuclidean(n1, n2)
	fmt.Println(gcd)

}

//approch 1 : min of two nums that divides both , else min--
//time complexity :  O(min(n1,n2))
func gcdofnums(n1 int, n2 int) int {

	return gcd(n1, n2, min(n1, n2))

}
func gcd(n1 int, n2 int, min int) int {
	if min == 1 {
		return min
	}

	if n1%min == 0 && n2%min == 0 {
		return min
	} else {
		return gcd(n1, n2, min-1)
	}
}

func min(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}

//Approch 2 : how we used to do in class
// Time :log(min(a,b))
func gcdEuclidean(n1 int, n2 int) int {

	if n2 == 0 {
		return n1
	}

	return gcdEuclidean(n2, n1%n2)

}
