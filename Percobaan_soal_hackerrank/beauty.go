package main

func main() {
	var i, j, k int32 = 20, 23, 6
	var res int32 = 0
	for i <= j {
		// fmt.Println(i, ",", reverse_int(i), "|", k)
		if (int(i)-reverse_int(i))%int(k) == 0 {
			res++
		}
		i++
	}
	// fmt.Println(res)
}

func reverse_int(n int32) int {
	new_int := 0
	for n > 0 {
		remainder := n % 10
		new_int *= 10
		new_int += int(remainder)
		n /= 10
	}
	return new_int
}
