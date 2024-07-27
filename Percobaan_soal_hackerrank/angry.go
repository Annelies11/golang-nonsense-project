package main

func main() {
	var k int32 = 3
	a := []int32{-1, -3, 4, 2}
	s := ""
	for _, num := range a {
		if num <= 0 {
			k--
		}
	}
	if k <= 0 {
		s = "YES"
	} else {
		s = "NO"
	}

}
