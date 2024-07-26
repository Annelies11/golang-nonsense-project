package main

import (
	"fmt"
)

func main() {
	h := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}
	word := "zaba"
	c := "abcdefghijklmnopqrstuvwxyz"
	var res, max int32 = 0, 0
	for _, car := range word {
		for j, cor := range c {
			if car == cor {
				res = h[j]
				if max < res {
					max = res
				}
			}
		}
	}
	res = max * int32(len(word))
	fmt.Println(res)
}
