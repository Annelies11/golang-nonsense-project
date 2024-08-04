package main

import (
	"fmt"
)

func main() {
	p := []int32{4, 3, 5, 1, 2}
	res := []int32{}
	for i := 1; i <= len(p); i++ {
		res = append(res, (FindIndex(FindIndex(int32(i), p)+1, p) + 1))
	}
	fmt.Println(res)
}

func FindIndex(n int32, arr []int32) int32 {
	var ret int32 = 0
	for i, num := range arr {
		if n == num {
			// break
			ret = int32(i)
		}
	}
	return ret
}
