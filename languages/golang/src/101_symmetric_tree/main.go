package _101_symmetric_tree

import "github.com/eliteGoblin/sky_ladder/common"

func isSymmetric(root *common.TreeNode) bool {
	if root == nil {
		return true
	}
	return isTwoTreeSym(root.Left, root.Right)
}

func isTwoTreeSym(tree1, tree2 *common.TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	} else if tree1 == nil || tree2 == nil {
		return false
	}
	if tree1.Val != tree2.Val {
		return false
	}
	return isTwoTreeSym(tree1.Left, tree2.Right) &&
		isTwoTreeSym(tree1.Right, tree2.Left)
}
