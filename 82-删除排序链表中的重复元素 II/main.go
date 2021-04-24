package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	start := &ListNode{0, head}
	curr := start
	for curr.Next != nil && curr.Next.Next != nil {
		if curr.Next.Val == curr.Next.Next.Val {
			val := curr.Next.Val
			for curr.Next != nil && curr.Next.Val == val {
				curr.Next = curr.Next.Next
			}
		} else {
			curr = curr.Next
		}
	}
	return start.Next
}
