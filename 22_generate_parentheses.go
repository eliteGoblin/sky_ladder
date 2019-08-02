package amazon

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	gen(n, n, "", &res)
	return res
}

func gen(left, right int, cur string, res *[]string) {
	if left < right || left < 0 || right < 0 {
		return
	}
	if left == 0 && right == 0 {
		*res = append(*res, cur)
	}
	gen(left-1, right, cur+"(", res)
	gen(left, right-1, cur+")", res)
}
