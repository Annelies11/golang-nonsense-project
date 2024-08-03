package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int32 = 1012
	var res int32 = 0
	for i, _ := range strconv.Itoa(int(n)) {
		d := string(strconv.Itoa(int(n))[i])
		a, _ := strconv.Atoi(d)
		if a == 0 {
			continue
		} else if n%int32(a) == 0 {
			res++
		}
	}
	// fmt.Println(len(strconv.Itoa(int(n))))
	// fmt.Println(string(strconv.Itoa(int(n))[0]))
	fmt.Println(res)
}
