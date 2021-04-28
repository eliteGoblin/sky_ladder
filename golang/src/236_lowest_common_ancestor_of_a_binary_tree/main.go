package twothreesix

import (
	"github.com/eliteGoblin/sky_ladder/common"
)

func lowestCommonAncestor(root, p, q *common.TreeNode) *common.TreeNode {
	pathP, _ := findPath(root, p)
	pathQ, _ := findPath(root, q)

	i := len(pathP) - 1
	j := len(pathQ) - 1
	for ; i >=0 && j >= 0;{
		if pathP[i] == pathQ[i] {
			return pathP[i]
		}
		i --
		j --
	}
	return root
}

func findPath(root, p *common.TreeNode)([]*common.TreeNode, bool) {
	if root == p {
		return []*common.TreeNode{root}, true
	}
	if root == nil {
		return nil, false
	}
	rc, ok := findPath(root.Left, p)
	if !ok {
		rc, ok = findPath(root.Right,p)
	}
	return append(rc, root), ok
}