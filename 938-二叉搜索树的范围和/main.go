package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
1. 二叉搜索树
	节点的左子树上的任何节点小于此节点
2. 当前节点值大于high，只考虑左子树
3. 当前节点值小于low，只考虑右子树
4. 介于中间，返回节点值与左子树的和与右子树的和
*/
// DFS
//func rangeSumBST(root *TreeNode, low int, high int) int {
//	if root == nil {
//		return 0
//	}
//	if root.Val > high {
//		return rangeSumBST(root.Left, low, high)
//	}
//	if root.Val < low {
//		return rangeSumBST(root.Right, low, high)
//	}
//	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
//}

// BFS
func rangeSumBST(root *TreeNode, low int, high int) int {
	queue := []*TreeNode{root}
	ans := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr == nil {
			continue
		}
		if curr.Val > high {
			queue = append(queue, curr.Left)
		} else if curr.Val < low {
			queue = append(queue, curr.Right)
		} else {
			queue = append(queue, curr.Left, curr.Right)
			ans += curr.Val
		}
	}
	return ans
}
