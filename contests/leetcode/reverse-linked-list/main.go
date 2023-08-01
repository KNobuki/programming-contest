package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	node := &ListNode{
		Val: head.Val,
	}
	for head.Next != nil {
		head = head.Next
		node = &ListNode{
			Val:  head.Val,
			Next: node,
		}
	}
	return node
}

func main() {

}
