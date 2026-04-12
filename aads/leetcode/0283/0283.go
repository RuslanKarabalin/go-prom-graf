package main

import "fmt"

func moveZeroes(nums []int) {
	n := len(nums)
	i := 0
	for i < n {
		if nums[i] == 0 {
			r := i + 1
			for r < n && nums[r] == 0 {
				r++
			}
			if r == n {
				return
			}
			nums[i], nums[r] = nums[r], nums[i]
		}
		i++
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
