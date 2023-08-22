package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	ret := &TreeNode{}
	var r1, r2, l1, l2 *TreeNode
	if root1 != nil {
		ret.Val += root1.Val
		l1 = root1.Left
		r1 = root1.Right
	}
	if root2 != nil {
		ret.Val += root2.Val
		l2 = root2.Left
		r2 = root2.Right
	}
	ret.Left = mergeTrees(l1, l2)
	ret.Right = mergeTrees(r1, r2)
	return ret
}
