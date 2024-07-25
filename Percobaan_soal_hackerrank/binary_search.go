package main

import (
	"fmt"
	"math"
)

func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	s := 16
	fmt.Println(FindMe(s, 0, len(num), num))

}

func FindMe(target, start, end int, arr []int) int {
	if start > end {
		return -1
	}
	middle := math.Floor((float64(start) + float64(end)) / 2)

	if arr[int(middle)] == target {
		return int(middle)
	}
	if arr[int(middle)] > target {
		return FindMe(target, start, int(middle)-1, arr)
	}
	if arr[int(middle)] < target {
		return FindMe(target, start, int(middle)+1, arr)
	}
	return 0
}
