package main

import (
	"fmt"
)

func main() {
	// var n, m, s int32 = 5, 2, 2
	// var n, m, s int32 = 7, 19, 2 // 6
	var n, m, s int32 = 3, 7, 3 // 3

	// var res int32 = 0

	var res int32 = 0
	res = s + m - 1
	res = res % n
	if res == 0 {
		fmt.Println(n)
	} else {
		fmt.Println(s - 1)
	}

}
