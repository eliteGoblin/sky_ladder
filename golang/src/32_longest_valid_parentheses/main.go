package longest_valid_parentheses

import (
	"github.com/eliteGoblin/sky_ladder/common"
)

func longestValidParentheses(s string) int {
	stk := common.NewStack(0)

	start := 0
	res := 0
	for i := range s {
		if s[i] == '(' {
			stk.Push(i)
		} else {
			if stk.Len() == 0 {
				start = i + 1
			} else {
				stk.Pop()
				var left int
				if stk.Len() > 0 {
					e, _ := stk.Top()
					left = e.(int) + 1
				} else {
					left = start
				}
				res = common.Max(res, i - left + 1)
			}
		}
	}
	return res
}