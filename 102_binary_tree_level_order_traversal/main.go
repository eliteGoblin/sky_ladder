package _102_binary_tree_level_order_traversal

import "github.com/eliteGoblin/sky_ladder/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *common.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	queue := make([]*common.TreeNode, 1)
	queue[0] = root

	for len(queue) > 0 {
		layer := make([]int, 0)
		next := make([]*common.TreeNode, 0)
		for _, v := range queue {
			layer = append(layer, v.Val)
			if v.Left != nil {
				next = append(next, v.Left)
			}
			if v.Right != nil {
				next = append(next, v.Right)
			}
		}
		res = append(res, layer)
		queue = next
	}
	return res
}
