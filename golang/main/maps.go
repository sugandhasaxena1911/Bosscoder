package main

import "fmt"

func main() {
	fmt.Println("NIL MAP")
	var m1 map[string]int // NIL MAP
	fmt.Println(m1)
	//m1["tina"] = 1 // gives error : assignment to entry in nil map

	m2 := make(map[string]int) // MAKE EMPTY MAP
	fmt.Println("m2  ", m2)

	m3 := make(map[string]int, 2) //MAKE EMPTY MAP
	fmt.Println("m3  ", m3)

	//
	var m6 = map[string]int{} // EMPTY MAP
	fmt.Println("m6  ", m6)
	m6["m6"] = 6

	m2["tina"] = 1
	m3["tina"] = 1
	fmt.Println(m2)
	fmt.Println(m3)

	//The initial capacity does not bound its size: maps grow to accommodate the number of items stored in them
	m3["sugandha"] = 2
	m3["manu"] = 3
	fmt.Println(len(m3))

	///EMPTY MAP
	var m4 = make(map[string]int, 5) // note the equality sign and not :=  , because  var is used
	fmt.Println("m4  ", m4)

	m4["tina"] = 1
	fmt.Println(m4)

	///EMPTY MAP
	var m5 = make(map[string]int)
	fmt.Println("m5  ", m5)

	m5["Hello"] = 8
	fmt.Println(m5)
}
