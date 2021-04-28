package _199_binary_tree_right_side_view

import "github.com/eliteGoblin/sky_ladder/common"

func rightSideView(root *common.TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return []int{}
	}
	layer := make([]*common.TreeNode, 1)
	layer[0] = root
	for len(layer) > 0 {
		res = append(res, layer[len(layer)-1].Val)
		nextLayer := make([]*common.TreeNode, 0)
		for _, v := range layer {
			if v.Left != nil {
				nextLayer = append(nextLayer, v.Left)
			}
			if v.Right != nil {
				nextLayer = append(nextLayer, v.Right)
			}
		}
		layer = nextLayer
	}
	return res
}
