package main

type ListNode struct {
	Next *ListNode
	Val  int
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var out *ListNode
	curOdd := out

	var even *ListNode
	curEven := even

	isOdd := true
	var i int
	for head != nil {
		if isOdd {
			if i == 0 {
				out = head
				curOdd = head
			} else {
				curOdd.Next = head
				curOdd = curOdd.Next
			}
		} else {
			if i == 1 {
				even = head
				curEven = head
			} else {
				curEven.Next = head
				curEven = curEven.Next
			}

		}

		isOdd = !isOdd
		head = head.Next
		i++
	}

	curOdd.Next = even
	if curEven != nil {
		curEven.Next = nil
	}

	return out
}
