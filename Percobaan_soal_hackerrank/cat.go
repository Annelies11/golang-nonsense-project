package main

import (
	"fmt"
)

func main() {
	// var x int32 = 1
	// var y int32 = 2
	// var z int32 = 3
	x := 1
	y := 3
	z := 2
	y = z - y
	if y < 0 {
		y = y * -1
	}

	x = z - y
	if x < 0 {
		x = x * -1
	}
	// fmt.Println(x, y)
	if x == y {
		fmt.Println("Mouse C")
	}
	if x < y {
		fmt.Println("Cat A")
	}
	if x > y {
		fmt.Println("Cat B")
	}
}
