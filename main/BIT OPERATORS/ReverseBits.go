package main

import "fmt"

func main() {
	var num uint32
	num = 139
	n := reverseBits(num)
	fmt.Printf("the number          %b\n", num)
	fmt.Printf("the reversed answer %b", n)

}

func reverseBits(num uint32) uint32 {
	var mask uint32
	var ans uint32
	for k := 0; k < 32; k++ {

		newpos := 0
		lastbit := (num >> k) & 1
		//move last bit to new position
		newpos = 31 - k
		mask = lastbit << (newpos)
		ans = ans | mask
	}

	return ans

}
