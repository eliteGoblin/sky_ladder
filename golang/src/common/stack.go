package common

import (
	"errors"
	"fmt"
)

type Stack []interface{}

func NewStack(cap int) *Stack {
	s := make(Stack, 0, cap)
	return &s
}

func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *Stack) Pop() error {
	if len(*s) == 0 {
		return errors.New("underflow")
	}
	*s = (*s)[:len(*s)-1]
	return nil
}

func (s *Stack) Top() (interface{}, error) {
	if len(*s) == 0 {
		return struct{}{}, errors.New("underflow")
	}
	return (*s)[len(*s)-1], nil
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Print() {
	fmt.Printf("Stack: %+v\n", *s)
}
