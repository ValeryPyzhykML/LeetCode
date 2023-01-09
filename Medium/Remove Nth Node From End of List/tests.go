package main

import (
	"fmt"
)

func arrayToLinkedList(array []int) *ListNode {
	head := &ListNode{}
	current := head
	for _, num := range array {
		current.Next = &ListNode{}
		current = current.Next
		current.Val = num
	}
	current.Next = nil
	return head.Next
}

func main() {
    test([]int{1,2,3,4,5}, 2, []int{1,2,3,5})
	test([]int{1,2,3,4,5}, 5, []int{2,3,4,5})
	test([]int{1,2}, 2, []int{2})
	test([]int{1,2}, 1, []int{1})
	test([]int{1}, 1, nil)
}

func test(array []int, n int, expectedResult []int) bool {
	actualResult := removeNthFromEnd(arrayToLinkedList(array),n);
	
	fmt.Println("==============================================================")
	fmt.Print(expectedResult)
	fmt.Println(" expected result")

	if expectedResult == nil {
		if actualResult == nil {
			fmt.Println("nil actual result")
			fmt.Println("OK")
			return true
		} else {
			fmt.Println("Not nil acctual result")
			fmt.Println("TEST FAILED: actual result must be nil")
			return true
		}
	}
	current := actualResult
	fmt.Printf("[")


	for i, expectedElement := range expectedResult {
		if current == nil {
			fmt.Printf("TEST FAILED: length of actural result (%d) != length of expected result (%d)\n", i + 1, len(expectedResult))
			return false
		}
		if expectedElement != current.Val {
			fmt.Printf("TEST FAILED: element i of actural result (%d) != element i of expected result (%d)\n", current.Val, expectedElement)
			return false
		}
		if (i == 0) {
			fmt.Printf("%d", current.Val)
		} else {
			fmt.Printf(" %d", current.Val)
		}
		
		current = current.Next
	}
	fmt.Println("] actual result")

	if current != nil {
		fmt.Printf("TEST FAILED: length of actural result > length of expected result (%d)\n", len(expectedResult))
		return false
	} 
	
	fmt.Println("OK")
	return true
}