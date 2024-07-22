package main

import (
	"fmt"
)

func main() {
	ar := []int{10, 20, 20, 10, 10, 30, 50, 10, 20}
	n := len(ar)
	res := 0
	for i := 0; i < n; i++ {
		j := i + 1
		for j < n {
			if ar[i] == ar[j] && ar[i] != 0 {
				ar[i], ar[j] = 0, 0
				res++
			}
			j++
		}
	}
	fmt.Println(res)

}
