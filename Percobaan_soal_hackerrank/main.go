package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 10
	s := string(strconv.FormatInt(int64(num), 2))
	res := 0
	fmt.Println(len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			res++
		}
	}
	fmt.Println(res)
}
