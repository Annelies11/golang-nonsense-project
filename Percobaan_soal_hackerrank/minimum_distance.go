package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int32{7, 1, 3, 4, 1, 7}
	var dis int32 = math.MaxInt32
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				dis = Min(dis, int32(j-i))
			}
		}
	}
	if dis == math.MaxInt32 {
		// return -1
		fmt.Println(-1)
	} else {
		// return dis
		fmt.Println(dis)
	}

}

func Min(a, b int32) int32 {
	if a > b {
		return b
	} else {
		return a
	}
}
