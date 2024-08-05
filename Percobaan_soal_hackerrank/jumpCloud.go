package main

import (
	"fmt"
)

func main() {
	c := []int32{0, 0, 0, 0, 1, 0} // 32
	var res int32 = 0
	i := 0
	for i < len(c)-1 {
		if i+2 < len(c) && c[i+2] == 0 {
			res++
			i += 2
		} else if i+1 < len(c) && c[i+1] == 0 {
			res++
			i++
		}
	}
	fmt.Println(res)
}
