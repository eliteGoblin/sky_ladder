package amazon

import "errors"

type ListNode struct {
	Val  int
	Next *ListNode
}

func max(a ...int) int {
	m := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > m {
			m = a[i]
		}
	}
	return m
}

func min(a ...int) int {
	m := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < m {
			m = a[i]
		}
	}
	return m
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
