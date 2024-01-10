package main

import (
	"fmt"
)

func main() {

	val := TurnKthBitON(135, 1)
	fmt.Println("Turn ON kth Bit")
	fmt.Printf("ANS  %b\n", val)

	fmt.Println("--------------------")
	val = TurnKthBitOFF(135, 0)
	fmt.Println("Turn OFF kth Bit")
	fmt.Printf("ANS  %b\n", val)

}

// bits positions start from 0 from right
func TurnKthBitON(n int, k int) int {
	mask := 1 << k
	fmt.Printf("NUM  %b\n", n)
	fmt.Printf("MASK %b\n", mask)
	return (n | mask)

}
func TurnKthBitOFF(n int, k int) int {
	mask := 1 << k
	fmt.Printf("NUM  %b\n", n)
	return (n & (^mask))

}
