package main

import (
	"fmt"
)

func main() {

	a := []int32{17, 28, 30}
	b := []int32{99, 16, 8}
	res := []int32{0, 0}
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			res[0]++
		} else {
			res[1]++
		}
	}
	fmt.Println(res)
}
