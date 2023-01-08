package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	current := head
	prev := current

	i := 0
	for current.Next != nil {
		if i >= n {
			prev = prev.Next
		}

		i++

		current = current.Next
	}

	if i == n-1 {
		return head.Next
	}

	if prev != nil {
		prev.Next = prev.Next.Next
	}

	return head
}