package common

import (
	"container/heap"
)

const (
	uint32Max = ^uint32(0)
	int32Max  = int32(uint32Max >> 1)
	int32Min  = -int32Max - 1
	uintMax   = ^uint(0)
	IntMax    = int(uintMax >> 1)
	IntMin    = -IntMax - 1
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

func (h *myHeap) Top() interface{} {
	if h.Len() > 0 {
		return h.data[0]
	}
	return nil
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

// first great than target
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

// first index that >= target

func lowerBoundIndex(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}

// 作用: 排序的数组中，找出重复元素的range: 第一个>=和第一个>
