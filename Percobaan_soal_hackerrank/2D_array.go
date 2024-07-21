package main

import (
	"fmt"
)

func main() {
	arr := [][]int32{
		{0, -4, -6, 0, -7, -6},
		{-1, -2, -6, -8, -3, -1},
		{-8, -4, -2, -8, -8, -6},
		{-3, -1, -2, -5, -7, -4},
		{-3, -5, -3, -6, -6, -6},
		{-3, -6, 0, -8, -6, -7}}
	// arr := [][]int32{{1, 1, 1, 0, 0, 0},
	// 	{0, 1, 0, 0, 0, 0},
	// 	{1, 1, 1, 0, 0, 0},
	// 	{0, 0, 2, 4, 4, 0},
	// 	{0, 0, 0, 2, 0, 0},
	// 	{0, 0, 1, 2, 4, 0}}
	var max, sum int32 = 0, 0
	for i := 2; i < len(arr); i++ {
		for j := 2; j < len(arr); j++ {
			// fmt.Println(arr[i-2][j-2], ",", arr[i-2][j-1], ",", arr[i-2][j])
			// fmt.Println("--", arr[i-2][j-1], "--")
			// fmt.Println(arr[i][j-2], ",", arr[i][j-1], ",", arr[i][j])
			// fmt.Println("____________")
			sum = arr[i-2][j-2] + arr[i-2][j-1] + arr[i-2][j] + arr[i-1][j-1] + arr[i][j-2] + arr[i][j-1] + arr[i][j]
			if sum > max {
				max = sum
			}
			sum = 0
		}
	}
	fmt.Println(max)

}
