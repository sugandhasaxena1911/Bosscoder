package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//fmt.Println(smallestGoodBase("1000000000000"))
	fmt.Println(smallestGoodBase64("13"))

}

// works for 32 bit
func smallestGoodBase(n string) string {
	no, e := strconv.Atoi(n)
	num := (no)
	if e != nil {
		fmt.Println(e)
		return ""
	}
	fmt.Println(num)

	for digits := 64; digits >= 2; digits-- {
		// b lies from 2 to n-1
		// binary search in arr
		low := (2)
		high := num - 1
		var answer int
		mid := (0)
		for low <= high {
			mid = (low + high) / 2
			// b^digits -1 / b-1
			r := math.Pow(float64(mid), float64(digits))
			answer = (int(r) - 1) / (mid - 1)
			if answer == num {
				return strconv.Itoa(int(mid))
			}
			if answer > num {
				high = mid - 1
			} else {
				low = mid + 1
			}

		}

	}
	return ""
}

func smallestGoodBase64(n string) string {
	num, _ := strconv.ParseInt(n, 10, 64)
	fmt.Println(num)
	var answer int64
	answer = num - 1

	for digits := (64); digits >= 2; digits-- {
		//fmt.Println(digits)
		// b lies from 2 to n-1
		// binary search in arr
		low := int64(2)
		high := num - 1

		var mid int64
		//var r int64
		for low <= high {
			mid = low + (high-low)/2
			fmt.Println("MID  ", mid, digits)
			// b^digits -1 / b-1
			//r = int64(math.Pow(float64(mid), float64(digits)))
			answer := CalculateNCompare((digits), mid, num)
			fmt.Println("ANSWER ", answer)

			if answer == 0 {
				return strconv.FormatInt(mid, 10)

			}
			//fmt.Println("r ", (r))

			//answer = ((r) - 1) / (mid - 1)

			/*if answer == num {
				return strconv.FormatInt(mid, 10)
			}*/
			if answer < 0 {
				low = mid + 1

			} else {
				high = mid - 1
			}

		}

	}
	return strconv.FormatInt(answer, 10)

}

// when base is right answer 0
// when is smaller -ve
// when base is grater 1+ve
func CalculateNCompare(digit int, base int64, num int64) int64 {
	if base > num {
		return 1
	}
	var mul int64
	mul = 1
	for i := 0; i < digit; i++ {
		num = num - mul
		if i != digit-1 && mul > num/base {
			return 1
		}
		if num < 0 {
			break
		}
		mul = mul * base
	}
	if num == 0 {
		return 0
	}
	if num < 0 {
		return 1
	} else {
		return -1
	}
}
