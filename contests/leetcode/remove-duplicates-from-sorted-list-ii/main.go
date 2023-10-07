package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	count := make(map[int]int)
	node := head
	for node != nil {
		count[node.Val]++
		node = node.Next
	}
	var ret, now *ListNode
	for head != nil {
		if count[head.Val] == 1 {
			if ret == nil {
				ret = &ListNode{Val: head.Val}
				now = ret
			} else {
				now.Next = &ListNode{Val: head.Val}
				now = now.Next
			}
		}
		head = head.Next
	}
	return ret
}
