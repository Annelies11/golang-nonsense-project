package main

import "fmt"

func main() {
	var n, res, sum int32 = 4, 5, 0
	for i := 1; i <= int(n); i++ {
		if i == 1 {
			// res = 5
			sum = res / 2
		} else {
			res = (res / 2) * 3
			sum = sum + (res / 2)
		}
		// fmt.Println(i, "||", res, "||", sum)
	}
	fmt.Println(sum)
}
