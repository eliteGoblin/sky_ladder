package amazon

import "container/heap"

type MedianFinder struct {
	left, right *myHeap
}

func Constructor() MedianFinder {
	dataLeft := make([]interface{}, 0, 512)
	dataRight := make([]interface{}, 0, 512)
	return MedianFinder{
		left: NewHeap(dataLeft, func(x, y int)bool{
			return x > y
		}),
		right: NewHeap(dataRight, func(x, y int)bool{
			return x < y
		}),
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.left, num)
	e := heap.Pop(this.left)
	heap.Push(this.right, e)
	if this.left.Len() < this.right.Len() {
		e := heap.Pop(this.right)
		heap.Push(this.left, e)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() > this.right.Len() {
		return float64(this.left.data[0].(int))
	}
	left := float64(this.left.data[0].(int))
	right := float64(this.right.data[0].(int))
	return (left + right) / 2.0
}
