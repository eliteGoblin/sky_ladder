package fourfourfive

import (
	"errors"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stk1 := NewStack(64)
	stk2 := NewStack(64)
	for p := l1; p != nil; p = l1.Next {
		stk1.Push(p)
	}
	for p := l2; p != nil; p = l2.Next {
		stk2.Push(p.Val)
	}
	var sum int
	var res *ListNode
	for stk1.Len() > 0 || stk2.Len() > 0 {
		if stk1.Len() > 0 {
			v, _ := stk1.Top()
			sum += v.(int)
			stk1.Pop()
		}
		if stk2.Len() > 0 {
			v, _ := stk2.Top()
			sum += v.(int)
			stk2.Pop()
		}
		res.Val = sum % 10
		head := &ListNode{
			Val: sum / 10,
			Next: res,
		}
		res = head
		sum /= 10
	}
	if res.Val == 0 {
		res = res.Next
	}
	return res
}

type stack []interface{}

func NewStack(cap int) *stack {
	s := make(stack, 0, cap)
	return &s
}

func (s *stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *stack) Pop() error {
	if len(*s) == 0 {
		return errors.New("underflow")
	}
	*s = (*s)[:len(*s)-1]
	return nil
}

func (s *stack) Top() (interface{}, error) {
	if len(*s) == 0 {
		return struct{}{}, errors.New("underflow")
	}
	return (*s)[len(*s)-1], nil
}

func (s *stack) Len() int {
	return len(*s)
}