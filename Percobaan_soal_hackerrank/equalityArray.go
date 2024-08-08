// s t k res
// qwerasdf qwerbsdf 2 no
// hackerhappy hackerrank 9 yes
// aba aba 7 yes
package main

import (
	"fmt"
)

func main() {
	arr := []int32{3, 3, 2, 1, 3}
	// arr := []int32{1, 2, 3, 1, 2, 3, 3, 3}

	m := make(map[int32]int32)
	var max, res int32 = 0, 0
	for _, num := range arr {
		max = num

		_, ok := m[num]
		// fmt.Println(num, "|", ok)
		if !ok {
			m[num] = 1
		} else {
			n := m[num]
			n++
			m[num] = n
		}
		if m[int32(max)] < m[num] {
			max = num
		}
	}
	for _, num := range m {
		res += num
	}
	res = res - m[max]
	delete(m, max)
	// fmt.Println(m)
	// fmt.Println(max)
	fmt.Println(res)
}
