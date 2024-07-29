package main

import (
	"fmt"
	"strconv"
)

func main() {
	// c := "abcdefgh"
	// r := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var n int = 0
	inp := ""
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&inp)
		cS := string(inp[0])
		rS, _ := strconv.Atoi(string(inp[1]))
		fmt.Println(cS, "||", rS)
	}
}
