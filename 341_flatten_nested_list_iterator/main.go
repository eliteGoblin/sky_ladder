package _341_flatten_nested_list_iterator

import "container/list"

type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool { return false }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return 0
}

// Set this NestedInteger to hold a single integer.
func (n NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []NestedInteger { return nil }

type NestedIterator struct {
	lst list.List
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	res := &NestedIterator{}
	res.lst.Init()
	for _, e := range nestedList {
		res.lst.PushBack(e)
	}
	return res
}

func (this *NestedIterator) Next() int {
	e := this.lst.Front()
	this.lst.Remove(e)
	intObj := e.Value.(*NestedInteger)
	return intObj.GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	for this.lst.Len() > 0 {
		e := this.lst.Front()
		intObj := e.Value.(*NestedInteger)
		if intObj.IsInteger() {
			return true
		}
		pre := e
		for _, v := range intObj.GetList() {
			newNode := this.lst.InsertAfter(v, pre)
			pre = newNode
		}
		this.lst.Remove(e)
	}
	return false
}
