package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str := "     oll123eH56     "
	// num := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	str = strings.Trim(str, " ")
	re := regexp.MustCompile(`\d`)
	str = re.ReplaceAllString(str, "")
	str = Reverse(str)
	fmt.Println(str)
}

func Reverse(s string) string {
	blok := []rune(s) //olleH
	for i, j := 0, len(blok)-1; i < j; i, j = i+1, j-1 {
		blok[i], blok[j] = blok[j], blok[i]
	}
	return string(blok)
}
//i = 0, j = 4; i < j; i++ j--
//olleH
//o,H = H,o
//Hlleo
//i = 1, j = 3
//l,e = e,l
//Hello
//i = 2, j = 2