package main

import "fmt"

func main() {
	// a := []int32{0, 1, 4}
	// n := a[0]
	var res int32 = 0
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			res++
		} else {
			res = res * 2
		}
		// fmt.Println(i, "==", res)
	}
	
}
