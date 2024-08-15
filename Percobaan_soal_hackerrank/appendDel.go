package main

import (
	"fmt"
)

func main() {
	s := "qwerasdf"
	t := "qwerbsdf"
	var k int32 = 6
	stop := 0
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			break
		}
		if i < len(t)-1 {
			break
		}
		stop++
	}
	del := len(s) - stop
	inc := len(t) - (len(s) - del)

	if inc == 0 && del <= int(k) {
		fmt.Println("Yes")
		//return "Yes"
	} else if del == 0 && (int(k)-inc)%2 != 0 {
		fmt.Println("No")
		//return "No"
	}

	op := del - inc
	if op <= int(k) {
		fmt.Println("Yes")
		//return "Yes"
	} else {
		fmt.Println("No")
		//return "No"
	}
}
