package main

import (
	"fmt"
)

func main() {
	a := []int32{1, 2, 3}
	queries := []int32{0, 1, 2}
	var k int32 = 2
	res := []int32{}
	n := len(a)
	k = k % int32(n)
	for i, _ := range queries {
		res = append(res, a[(int32(n)-k+int32(i))%int32(n)])
	}
	fmt.Println(res)
}

// func GakEfektif( a []int32, k int32){
// 	for i := 0; i < int(k); i++ {
// 		j := len(a)
// 		temp := a[j-1]
// 		// fmt.Println(temp)
// 		for {
// 			a[j-1] = a[j-2]
// 			j--
// 			if j-2 == 0 {
// 				a[1] = a[0]
// 				a[0] = temp
// 				break
// 			}
// 		}
// 		// fmt.Println(a)
// 	}
// 	for i, num := range queries {
// 		queries[i] = a[num]
// 	}
// 	fmt.Println(queries)

// }
