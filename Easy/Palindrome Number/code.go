package main

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 { return false }

	nums := strconv.Itoa(x)
	center := len(nums) / 2
	last := len(nums) - 1
	for i := 0; i < center; i++ {
		if nums[i] != nums[last-i] { return false }
	}
	return true
}
