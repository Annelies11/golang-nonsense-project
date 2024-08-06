package main

import "fmt"

// Pelajari lagi
func main() {
	arr := []int32{1, 2, 4, 5, 7, 8, 10} // 3
	var d int32 = 3
	// var a, b, c int32 = arr[0], 0, 0
	var res int32 = 0
	m := make(map[int]int)
	c := 0
	for _, num := range arr {
		m[int(num)] = c + 1
	}
	for _, num := range arr {
		a := m[int(num)+int(d)]
		b := m[int(num)+int(d)+int(d)]
		if checkExist(int32(a), arr) && checkExist(int32(b), arr) {
			res++
		}
	}
	fmt.Println(res)
}

func checkExist(n int32, arr []int32) bool {
	ret := false
	for _, item := range arr {
		if n == item {
			ret = true
		}
	}
	return ret
}
