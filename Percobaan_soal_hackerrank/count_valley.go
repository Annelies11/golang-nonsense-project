package main

import "fmt"

func main() {
	//steps int32, path string
	var path string = "DDUUDDUDUUUD"
	cont := 0
	res := 0
	for i := 0; i < len(path); i++ {
		prev := cont
		if path[i] == 'D' {
			cont--
		} else {
			cont++
		}
		if prev == -1 && cont == 0 {
			res++
		}
	}
	// fmt.Printf("%q", path[0])
	// fmt.Println(len(path))
	fmt.Println(res)
}
