package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
二叉树的中序遍历
(中，前，后)指的是节点值所居的位置
中序遍历 = 左子树最前，值居中，右子树最后
*/
func increasingBST(root *TreeNode) *TreeNode {
	vals := []int{}
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node != nil {
			inorder(node.Left)
			vals = append(vals, node.Val)
			inorder((node.Right))
		}
	}
	inorder(root)

	dummyNode := &TreeNode{}
	currNode := dummyNode
	for _, v := range vals {
		currNode.Right = &TreeNode{Val: v}
		currNode = currNode.Right
	}
	return dummyNode.Right
}
