package main

import "fmt"

func main() {
	test([]int{2, 7, 11, 15}, 9, []int{0,1})
	test([]int{2, 11, 15, 6}, 8, []int{0,3})
	test([]int{3, 2, 4}, 6, []int{1,2})
	test([]int{1, 1, 3, 2, 4}, 6, []int{3,4})
	test([]int{1, 1, 3, 1, 4}, 6, nil)
}

func test(nums []int, target int, expectedResult []int) {
	result := twoSum(nums, target)
	if (result == nil && expectedResult == nil) || (result[0] == expectedResult[0] && result[1] == expectedResult[1]) {
		fmt.Printf("OK result %v == %v expected result \n", result, expectedResult)
	} else {
		fmt.Printf("ERROR result %v != %v expected result \n", result, expectedResult)
	}
}