package main

import "fmt"

// Perlu dipahami lebih lanjut
func main() {
	c := []int32{0, 0, 1, 0, 0, 1, 1, 0} //92 1 1 1 0 1 1 0 0 0 0
	// c := []int32{1, 1, 1, 0, 1, 1, 0, 0, 0, 0} //92
	var k int32 = 2
	var e int32 = 100
	i := 0
	for {
		e = e - 1 - 2*c[i]
		i = (i + int(k)) % len(c)
		if i == 0 {
			break
		}
	}

	fmt.Println(e)

}
