package amazon

var mp = map[byte]byte{
	'}': '{',
	']': '[',
	')': '(',
}

func isValid(s string) bool {
	stk := NewStack(len(s))
	bytes := []byte(s)
	for _, v := range bytes {
		switch v {
		case '{':
			fallthrough
		case '[':
			fallthrough
		case '(':
			stk.Push(v)
		default:
			top, err := stk.Top()
			if err != nil ||
				mp[v] != top.(byte) {
				return false
			}
			stk.Pop()
		}
	}
	return stk.Len() == 0
}
