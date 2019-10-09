package A

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	level := make([]*TreeNode, 1)
	level[0] = root
	leftToRight := true
	for len(level) > 0 {
		nextLevel := make([]*TreeNode, 0)
		value := make([]int, 0)
		for i := range level {
			if level[i].Left != nil {
				nextLevel = append(nextLevel, level[i].Left)
			}
			if level[i].Right != nil {
				nextLevel = append(nextLevel, level[i].Right)
			}
		}
		start, end, delta := 0, len(level), 1
		if !leftToRight {
			start, end, delta = len(level)-1, -1, -1
		}
		for ; start != end; start += delta {
			value = append(value, level[start].Val)
		}
		res = append(res, value)
		leftToRight = !leftToRight
		level = nextLevel
	}
	return res
}
