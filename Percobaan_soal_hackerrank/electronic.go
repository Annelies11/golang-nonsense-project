package main

import (
	"fmt"
)

func main() {
	var b int32 = 5
	keyboards := []int32{4}
	drives := []int32{5}
	var res int32 = -1
	var prev int32 = -1

	for _, num := range keyboards {
		for _, nom := range drives {
			if res < num+nom {
				res = num + nom
			}
			if res > b {
				res = prev
			} else {
				prev = res
			}
		}
	}

	fmt.Println(res)
	res += 1
	fmt.Println(res)
}
