package amazon

type MinStack struct {
	value *stack
	min   *stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		value: NewStack(512),
		min:   NewStack(512),
	}
}

func (this *MinStack) Push(x int) {
	if this.min.Len() > 0 {
		v, _ := this.min.Top()
		if x <= v.(int) {
			this.min.Push(x)
		}
	} else {
		this.min.Push(x)
	}
	this.value.Push(x)
}

func (this *MinStack) Pop() {
	v, _ := this.value.Top()
	m, _ := this.min.Top()
	if v.(int) == m.(int) {
		this.min.Pop()
	}
	this.value.Pop()
}

func (this *MinStack) Top() int {
	v, _ := this.value.Top()
	return v.(int)
}

func (this *MinStack) GetMin() int {
	v, _ := this.min.Top()
	return v.(int)
}
