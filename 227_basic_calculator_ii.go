package A

// 1 + 1
// 1 + 2 * 3
func calculate(s string) int {
	var num int
	var op byte

	stk := NewStack(521)
	op = '+'

	for i := range s {
		if s[i] >= '0' {
			num = num*10 + int(s[i]-'0')
		}
		if s[i] < '0' && s[i] != ' ' || i == len(s)-1 {
			switch op {
			case '+':
				stk.Push(num)
			case '-':
				stk.Push(-num)
			case '*':
				t, _ := stk.Top()
				stk.Pop()
				stk.Push(t.(int) * num)
			case '/':
				t, _ := stk.Top()
				stk.Pop()
				stk.Push(t.(int) / num)
			}
			op = s[i]
			num = 0
		}
	}
	res := 0
	for stk.Len() > 0 {
		t, _ := stk.Top()
		res += t.(int)
		stk.Pop()
	}
	return res
}
