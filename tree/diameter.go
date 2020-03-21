package tree

var ans int

func DiameterOfBinaryTree(root *TreeNode) int {
	ans = 1

	depth(root)
	return ans - 1
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := depth(root.Left)
	right := depth(root.Right)

	ans = max(ans, left + right + 1)
	return max(left,right) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}