package _543_diameter_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	maxLenPassNode(&res, root)
	return res - 1
}

func maxLenPassNode(maxDiam *int, root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxLenPassNode(maxDiam, root.Left)
	r := maxLenPassNode(maxDiam, root.Right)
	current := l + r + 1
	if current > *maxDiam {
		*maxDiam = current
	}
	if l >= r {
		return l + 1
	} else {
		return r + 1
	}
}
