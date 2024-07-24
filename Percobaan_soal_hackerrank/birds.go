package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}
	sum := []int{0, 0, 0, 0, 0}
	// m := []int{}
	res := 0
	for i := 1; i <= 5; i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j] == i {
				sum[i-1]++
			}
		}
	}
	k := 0
	for true {
		if sum[k] == Max(sum) {
			res = k + 1
			break
		} else {
			k++
		}
	}
	fmt.Println(res)
}

func Max(data []int) int {
	max := data[0]
	for _, num := range data {
		if num > max {
			max = num
		}
	}
	return max
}
