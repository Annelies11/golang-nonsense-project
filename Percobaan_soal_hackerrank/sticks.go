package main

import "fmt"

func main() {
	arr := []int32{5, 4, 4, 2, 2, 8}
	res := []int32{}
	for {
		s := Min(arr)
		j := 0
		i := 0
		fmt.Println(arr)
		for k, num := range arr {
			if num > 0 {
				arr[k] -= s
				j++
			} else {
				i++
			}
		}
		res = append(res, int32(j))
		if i == len(arr)-1 {
			break
		}
	}

	fmt.Println(res)
}

func Min(arr []int32) int32 {
	min := arr[0]
	for _, num := range arr {
		if num < min && num > 0 {
			min = num
		}
	}
	return min
}
