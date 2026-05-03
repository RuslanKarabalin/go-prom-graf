package main

import (
	"fmt"
	"math"
)

func findMaxConsecutiveOnes(nums []int) int {
	n := len(nums)
	m := math.MinInt
	l := 0
	r := 0
	for r < n {
		if nums[r] == 1 {
			r++
		} else {
			m = max(m, r-l)
			for nums[r] == 0 {
				if r+1 < n {
					r++
				} else {
					return m
				}
			}
			l = r
		}
	}
	m = max(m, r-l)
	return m
}

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	res := findMaxConsecutiveOnes(nums)
	fmt.Println(res)
}
