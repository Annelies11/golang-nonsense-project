package main

import (
	"fmt"
)

// 2 7 1014
// 1 1 1015
func main() {
	d1, m1, y1 := 2, 7, 2015
	d2, m2, y2 := 1, 2, 2014
	var res int32 = 0
	if y2 < y1 {
		res = 10000
		// fine = false
	} else if m2 < m1 {
		res += int32((m1 - m2) * 500)
		// fine = false
	} else if d2 < d1 {
		res += int32((d1 - d2) * 15)
	}
	fmt.Println(res)

}
