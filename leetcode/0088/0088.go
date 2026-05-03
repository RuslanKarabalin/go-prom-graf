package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, 0, m+n)
	l, k := 0, 0

	for l < m && k < n {
		f := nums1[l]
		s := nums2[k]
		if f < s {
			res = append(res, f)
			l++
		} else {
			res = append(res, s)
			k++
		}
	}
	for l < m {
		f := nums1[l]
		res = append(res, f)
		l++
	}
	for k < n {
		f := nums2[k]
		res = append(res, f)
		k++
	}
	copy(nums1, res)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Printf("%v\n", nums1)
}
