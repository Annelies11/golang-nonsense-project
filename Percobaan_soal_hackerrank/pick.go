package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int32{1, 2, 2, 3, 1, 2} // 5 {1,2,2,1,2}
	// a := []int32{1, 1, 2, 2, 4, 4, 5, 5, 5} //5 {4,4,5,5,5}
	// a := []int32{4, 6, 5, 3, 3, 1} //3 {4,3,3}
	same := []int32{0, 0}
	var res, prev int32 = 0, 0
	for i, num := range a {
		same[0] = num
		for j, nom := range a {
			if i != j {
				fmt.Println(num, "-", nom, "= abs ", math.Abs(float64(num)-float64(nom)))
			}
			if math.Abs(float64(num)-float64(nom)) <= 1 {

				if nom != same[0] {
					same[1] = nom
					res++
				}
				if nom == same[0] || nom == same[1] {
					res++
				}
			}
			if prev < res {
				prev = res
			}
			res = 0
			same[1] = 0
		}
	}
	res = prev
	fmt.Println(res)
}
