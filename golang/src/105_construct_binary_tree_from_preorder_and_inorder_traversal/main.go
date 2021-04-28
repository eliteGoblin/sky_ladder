package _105_construct_binary_tree_from_preorder_and_inorder_traversal

import "github.com/eliteGoblin/sky_ladder/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *common.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &common.TreeNode{
		Val: preorder[0],
	}
	leftLen := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			leftLen = i
			break
		}
	}
	root.Left = buildTree(preorder[1:1+leftLen], inorder[:leftLen])
	root.Right = buildTree(preorder[1+leftLen:], inorder[1+leftLen:])
	return root
}

// 最费劲的是搞定inorder和preorder的index: 递归不用太考虑右半边，精髓在left; 右边的preorder和inorder起止肯定一样
