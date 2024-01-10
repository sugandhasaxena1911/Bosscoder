package main

import "fmt"

func main() {
	fmt.Println("INT,FLOAT, STRING , ARRAY,STRUCT , BOOL , RUNE BYTE is PASS BY VALUE")
	fmt.Println("SLICES , MAPS is PASS BY REFERENCE")

	a := 1
	checkint(a)
	fmt.Println(a)

	b := 1.5
	checkfloat(b)
	fmt.Println(b)

	c := "sugandha"
	checkstring(c)
	fmt.Println(c)

	cc := "sugandha"
	checkstringV2(cc)
	fmt.Println(cc)

	d := []int{1, 2, 3, 4, 5, 6}
	checkslice(d)
	fmt.Println(d)

	e := []int{1, 2, 3, 4, 5, 6}
	checkslicev2(e)
	fmt.Println(e)

	f := []int{1, 2, 3, 4, 5, 6}
	checkslicev3(f)
	fmt.Println(f)

	fnew := f[0:3] /// DOESNOT CREATE A NEW UNDERLYING ARRAY
	fnew[0] = 999
	fmt.Println("AFTER f & fnew ", f, fnew)

	g := map[string]int{"tina": 0}
	checkmap(g)
	fmt.Println(g)

}

func checkint(a int) {
	a = 10
}

func checkfloat(a float64) {
	a = 10.6

}

func checkstring(a string) {
	a = "sugandha saxena"
}

func checkstringV2(a string) {
	a = a + "saxena"
}

func checkslice(a []int) {
	a = []int{7, 8, 9, 10} // creates new underlying array
}

func checkslicev2(a []int) {
	a = append(a, 7, 8, 9, 10) // creates new underlying array

}

func checkslicev3(a []int) {
	a[1] = 101
}

func checkmap(a map[string]int) {
	a["sugandha"] = 1

}
