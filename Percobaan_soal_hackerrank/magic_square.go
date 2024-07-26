package main

import (
	"fmt"
	"math"
)

func main() {

	// s := [][]int32{{4, 9, 2},
	// 	{3, 5, 7},
	// 	{8, 1, 5}}
	// s := [][]int32{{4, 8, 2},
	// 	{4, 5, 7},
	// 	{6, 1, 6}}
	s := [][]int32{{2, 5, 4},
		{4, 6, 9},
		{4, 5, 2}}

	mag := [][]int32{{8, 1, 6, 3, 5, 7, 4, 9, 2}, {6, 1, 8, 7, 5, 3, 2, 9, 4}, {4, 3, 8, 9, 5, 1, 2, 7, 6}, {2, 7, 6, 9, 5, 1, 4, 3, 8}, {2, 9, 4, 7, 5, 3, 6, 1, 8}, {4, 9, 2, 3, 5, 7, 8, 1, 6}, {6, 7, 2, 1, 5, 9, 8, 3, 4}, {8, 3, 4, 1, 5, 9, 6, 7, 2}}

	var res, min int32 = 0, 100
	temp := []int32{}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			temp = append(temp, s[i][j])
		}
	}

	for i := 0; i < len(mag); i++ {
		for j := 0; j < len(mag[0]); j++ {
			if mag[i][j] != temp[j] && res <= min {
				res += int32(math.Abs(float64(mag[i][j]) - float64(temp[j])))
			}
			if res >= min {
				break
			}
		}
		if res < min {
			min = res
		}
		res = 0
	}

	fmt.Println(min)
}
