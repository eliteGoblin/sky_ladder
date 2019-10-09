package amazon

import (
	"container/heap"
	"errors"
)

const (
	uint32Max = ^uint32(0)
	int32Max  = int32(uint32Max >> 1)
	int32Min  = -int32Max - 1
	uintMax   = ^uint(0)
	intMax    = int(uintMax >> 1)
	intMin    = -intMax - 1
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

type myHeap struct {
	data      []interface{}
	predicate func(x, y interface{}) bool
}

func NewHeap(data []interface{}, predicate func(x, y interface{}) bool) *myHeap {
	hp := &myHeap{
		data:      data,
		predicate: predicate,
	}
	heap.Init(hp)
	return hp
}

func (h myHeap) Len() int {
	return len(h.data)
}

func (h myHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h myHeap) Less(i, j int) bool {
	return h.predicate(h.data[i], h.data[j])
}

func (h *myHeap) Push(x interface{}) {
	h.data = append(h.data, x)
}

func (h *myHeap) Pop() interface{} {
	len := len(h.data)
	x := h.data[len-1]
	h.data = h.data[:len-1]
	return x
}

// 值得注意的
// 区分heap.Push和自己实现的Push
// predicate, tricky, 必须每次传递data,否则闭包还是调用之前的slice结构，slice伸缩了，predicate还是不知道;
// predicate最好传递值，而不是index

func upperBoundIndex(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}
