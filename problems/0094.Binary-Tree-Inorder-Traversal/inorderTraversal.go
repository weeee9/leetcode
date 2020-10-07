package leetcode

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	inorder(root, &result)
	return result
}

func inorder(node *TreeNode, arr *[]int) {
	if node != nil {
		inorder(node.Left, arr)
		*arr = append(*arr, node.Val)
		inorder(node.Right, arr)
	}
}
