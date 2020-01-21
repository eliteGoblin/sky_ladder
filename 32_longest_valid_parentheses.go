package amazon

func longestValidParentheses(s string) int {
	left, right := 0, 0
	result := 0
	for i := range s {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			result = max(result, left)
		} else if left < right {
			left, right = 0, 0
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			result = max(result, left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return result * 2
}

func longestValidParentheses2(s string) int {
	stk := NewStack(len(s) / 2)
	result := 0
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stk.Push(i)
		} else {
			if stk.Len() == 0 {
				start = i + 1
			} else {
				stk.Pop()
				if stk.Len() > 0 {
					topE, _ := stk.Top()
					result = max(result, i-topE.(int))
				} else {
					result = max(result, i-start+1)
				}
			}
		}
	}
	return result
}


func longestValidParentheses(s string) int {
	stk := NewStack()
}
