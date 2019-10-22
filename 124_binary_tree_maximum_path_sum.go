package amazon

func maxPathSum(root *TreeNode) int {
	res := intMin
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left := max(dfs(root.Left, res), 0)
	right := max(dfs(root.Right, res), 0)
	*res = max(*res, left+right+root.Val)
	return root.Val + max(left, right)
}
