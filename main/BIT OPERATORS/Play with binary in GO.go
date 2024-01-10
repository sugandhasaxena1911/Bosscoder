package main

import "fmt"

func main() {
	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	//var c uint = 0

	c := a & b /* 12 = 0000 1100 */
	fmt.Println(a)
	fmt.Printf("Binary %b ", a)
	fmt.Printf("\nType of a variable %T ", a)

	fmt.Printf("\nLine 1 - Value of c is %d\n", c)
	fmt.Println("Line 1 - Value of c is", c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("Line 2 - Value of c is %d\n", c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("Line 3 - Value of c is %d\n", c)

	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("Line 4 - Value of c is %d\n", c)

	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Printf("Line 5 - Value of c is %d\n", c)
}
