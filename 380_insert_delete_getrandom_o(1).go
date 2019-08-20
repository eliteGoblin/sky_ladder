package amazon

import "math/rand"

type RandomizedSet struct {
	data []int
	dict map[int]int
}

func Constructor380() RandomizedSet {
	return RandomizedSet{
		data: make([]int, 0, 1024),
		dict: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.dict[val]; ok {
		return false
	}
	this.data = append(this.data, val)
	this.dict[val] = len(this.data) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if pos, ok := this.dict[val]; !ok {
		return false
	} else {

		// edge case: pos == lastPos, 或者只有一个元素，或者删除的是最后一个
		last := this.data[len(this.data)-1]
		this.dict[last] = this.dict[val]
		this.data[pos] = last
		this.data = this.data[:len(this.data)-1]
		delete(this.dict, val) // 必须放在最后: edge case: pos == lastpos, 先删除后insert会在map残留

		return true
	}
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[rand.Intn(len(this.data))]
}
