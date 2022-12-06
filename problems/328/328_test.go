package main

import (
	"testing"
)

func Test328(t *testing.T) {
	for _, tc := range []struct {
		in   *ListNode
		want *ListNode
	}{
		{
			in:   newListNode(1, 2, 3, 4, 5),
			want: newListNode(1, 3, 5, 2, 4),
		},
		{
			in:   newListNode(2, 1, 3, 5, 6, 4, 7),
			want: newListNode(2, 3, 6, 7, 1, 5, 4),
		},
	} {
		want := make([]int, 0)
		for head := tc.want; ; {
			if head == nil {
				break
			}
			want = append(want, head.Val)
			head = head.Next
		}

		res := oddEvenList(tc.in)
		got := make([]int, 0)
		for head := res; ; {
			if head == nil {
				break
			}
			got = append(got, head.Val)
			head = head.Next
		}

		if !slicesAreEqual(want, got) {
			t.Errorf("\ninvalid result\nwant: %#v\ngot: %#v", want, got)
		}
	}
}

func slicesAreEqual[E comparable](s1, s2 []E) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func newListNode(nums ...int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	out := new(ListNode)
	cur := out
	for i, n := range nums {
		if i == 0 {
			cur.Val = n
			continue
		}

		cur.Next = &ListNode{Val: n}
		cur = cur.Next
	}
	return out
}
