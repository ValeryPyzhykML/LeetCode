package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	start := &ListNode{Next: head}
	current := &start

	for n != 0 {
		n--
		current = current.Next
	}

	prev := current

	if prev != nil {
		prev.Next = prev.Next.Next
	}

	return start.Next
}