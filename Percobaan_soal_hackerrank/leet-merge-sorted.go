package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	// for i, j := m, 0; j < n; i, j = i+1, j+1 {
	// 	nums1[i] = nums2[j]
	// }
	for i, j := (m+n)-1, n-1; i > m-1; i, j = i-1, j-1 {
		nums1[i] = nums2[j]
	}
	fmt.Println(nums1)

	sort.Ints(nums1)

	fmt.Println(nums1)
}
