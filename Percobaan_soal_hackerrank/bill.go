package main

import (
	"fmt"
)

func main() {
	bill := []int32{3, 10, 2, 9}
	k := 1
	b := 12
	sum := 0

	for i, num := range bill {
		if i != k {
			sum += int(num)
		}
	}
	sum = sum / 2
	sum = b - sum
	if sum == 0 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(sum)
	}

	//k = 1 b = 12 print 5
	//k = 1 b = 7 print Bon Appetit
}
