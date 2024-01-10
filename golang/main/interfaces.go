package main

import (
	"fmt"
)

func assert(i interface{}) {
	//s := i.(int) //get the underlying int value from i
	//fmt.Println(s)

	v, ok := i.(int)
	fmt.Println("interface v, Ok ", v, ok)

	//v, ok = i.(type)    OUTSIDE SWITCH NOT ALLOWED
	fmt.Println(v, ok)

}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	var s interface{} = 56
	fmt.Printf("TYPE is %T ", s)
	fmt.Println("")
	assert(s)

	var s1 interface{} = "Steven Paul"
	fmt.Printf("TYPE is %T ", s)
	assert(s1)

	findType("Naveen")
	findType(77)
	findType(89.98)
}
