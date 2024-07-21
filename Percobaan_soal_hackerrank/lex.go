package main

import (
	"fmt"
)

func main() {
	var s string = "abcd"
	var t string = "abdc"
	// byteArray := []byte(s)
	fmt.Println(SumByte(s))
	fmt.Println(SumByte(t))
}

func SumByte(data string) int {
	sum := 0
	for i := 0; i < len([]byte(data)); i++ {
		sum += int([]byte(data)[i])
	}
	return sum
}
