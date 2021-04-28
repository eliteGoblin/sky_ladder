package _173_binary_search_tree_iterator

import "github.com/eliteGoblin/sky_ladder/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	stk *common.Stack
}

func Constructor(root *common.TreeNode) BSTIterator {
	stk := common.NewStack(128)
	leftMost(stk, root)
	return BSTIterator{
		stk: stk,
	}
}

func leftMost(stk *common.Stack, root *common.TreeNode) {
	for root != nil {
		stk.Push(root)
		root = root.Left
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	topNode, _ := this.stk.Top()
	top := topNode.(*common.TreeNode)
	this.stk.Pop()
	if top.Right != nil {
		leftMost(this.stk, top.Right)
	}
	return top.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.stk.Len() > 0
}

// 非递归中序遍历BST:  中序遍历BST即为有序遍历
