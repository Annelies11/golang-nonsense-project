package main

import (
	"fmt"
)

func main() {
	var n int32 = 15603
	var p int32 = 6957 //5
	//15603
	// 6957
	//output : 3478
	var res int32 = 0
	// 1 | 4
	// 4 | 1
	if p-1 < n-p {
		for i := 0; i < int(p); i++ {
			if i%2 == 0 && i != 0 {
				res++
			}
		}
	} else {
		for i := int(n); i > int(p); {
			if i%2 == 0 {
				res++
			}
		}
	}
	fmt.Println(res)

}
