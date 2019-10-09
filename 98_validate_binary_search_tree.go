package amazon

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValid(root, intMin, intMax)
}

func isValid(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return isValid(root.Left, lower, root.Val) &&
		isValid(root.Right, root.Val, upper)
}
