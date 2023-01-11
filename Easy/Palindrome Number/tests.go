package main

import "fmt"

func main() {
	test(-121, false)
	
	test(121, true)
	test(22229877892222, true)

	test(123, false)
	test(3123, false)
}

func test(num int, expectedResult bool) {
	result := isPalindrome(num)
	if (result == expectedResult) {
		fmt.Printf("OK result %v == %v expected result \n", result, expectedResult)
	} else {
		fmt.Printf("ERROR result %v != %v expected result \n", result, expectedResult)
	}
}