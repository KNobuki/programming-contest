package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	vis := make(map[*ListNode]bool)
	for head != nil {
		if vis[head] {
			return false
		}
		vis[head] = true
		head = head.Next
	}
	return true
}
