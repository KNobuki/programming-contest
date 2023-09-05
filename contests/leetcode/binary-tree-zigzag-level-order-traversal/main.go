package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func zigzagLevelOrder(root *TreeNode) [][]int {
//	type nodeLevel struct {
//		node  *TreeNode
//		level int
//	}
//	que := []nodeLevel{
//		{node: root, level: 0},
//	}
//	var ret [][]int
//	for len(que) > 0 {
//		cur := que[0]
//		que = que[1:]
//		if cur.node == nil {
//			continue
//		}
//		if len(ret) <= cur.level {
//			ret = append(ret, []int{})
//		}
//		if cur.level%2 == 0 {
//			ret[cur.level] = append(ret[cur.level], cur.node.Val)
//		} else {
//			ret[cur.level] = append([]int{cur.node.Val}, ret[cur.level]...)
//		}
//		que = append(que, nodeLevel{node: cur.node.Left, level: cur.level + 1}, nodeLevel{node: cur.node.Right, level: cur.level + 1})
//	}
//	return ret
//}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var ret [][]int
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(ret) <= level {
			ret = append(ret, []int{})
		}
		if level%2 == 0 {
			ret[level] = append(ret[level], node.Val)
		} else {
			ret[level] = append([]int{node.Val}, ret[level]...)
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)
	return ret
}
