package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pParents := make([]*TreeNode, 0)
	qParents := make([]*TreeNode, 0)
	var pDfs func(now *TreeNode) bool
	pDfs = func(now *TreeNode) bool {
		if now == nil {
			return false
		}
		pParents = append(pParents, now)
		if now.Val == p.Val {
			return true
		}
		if pDfs(now.Left) {
			return true
		}
		if pDfs(now.Right) {
			return true
		}
		pParents = pParents[:len(pParents)-1]
		return false
	}
	pDfs(root)
	var qDfs func(now *TreeNode) bool
	qDfs = func(now *TreeNode) bool {
		if now == nil {
			return false
		}
		qParents = append(qParents, now)
		if now.Val == q.Val {
			return true
		}
		if qDfs(now.Left) {
			return true
		}
		if qDfs(now.Right) {
			return true
		}
		qParents = qParents[:len(qParents)-1]
		return false
	}
	qDfs(root)
	for i := 0; i < len(pParents) && i < len(qParents); i++ {
		if pParents[i].Val != qParents[i].Val {
			return pParents[i-1]
		}
	}
	if len(pParents) > len(qParents) {
		return q
	}
	return p
}
