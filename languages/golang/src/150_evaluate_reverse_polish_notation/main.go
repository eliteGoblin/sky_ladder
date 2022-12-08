package _150_evaluate_reverse_polish_notation

import (
	"strconv"

	"github.com/eliteGoblin/sky_ladder/common"
)

func evalRPN(tokens []string) int {
	stk := common.NewStack(len(tokens))
	for _, v := range tokens {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			op2, _ := stk.Top()
			stk.Pop()
			op1, _ := stk.Top()
			stk.Pop()
			i1, i2 := op1.(int), op2.(int)
			if v == "+" {
				stk.Push(i1 + i2)
			} else if v == "-" {
				stk.Push(i1 - i2)
			} else if v == "*" {
				stk.Push(i1 * i2)
			} else {
				stk.Push(i1 / i2)
			}
		} else {
			i, _ := strconv.Atoi(v)
			stk.Push(i)
		}
	}
	res, _ := stk.Top()
	return res.(int)
}
