package amazon

import "container/list"

type value struct {
	v    int
	freq int
}

type LFUCache struct {
	capacity int
	m        map[int]value
	freq     map[int]*list.List
	ptr      map[int]*list.Element
	minFreq  int
}

func ConstructorLFU(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		m:        make(map[int]value),
		freq:     make(map[int]*list.List),
		ptr:      make(map[int]*list.Element),
		minFreq:  0,
	}
}

func (this *LFUCache) Get(key int) int {
	if _, ok := this.m[key]; !ok {
		return -1
	}
	this.freq[this.m[key].freq].Remove(this.ptr[key])
	o := this.m[key]
	o.freq++
	this.m[key] = o
	lst := this.freq[this.m[key].freq]
	if lst == nil {
		this.freq[this.m[key].freq] = list.New()
	}
	this.freq[this.m[key].freq].PushBack(key)
	this.ptr[key] = this.freq[this.m[key].freq].Back()
	if this.freq[this.minFreq].Len() == 0 {
		this.minFreq++
	}
	return this.m[key].v
}

func (this *LFUCache) Put(key int, v int) {
	if this.capacity <= 0 {
		return
	}
	if this.Get(key) != -1 {
		o := this.m[key]
		o.v = v
		this.m[key] = o
		return
	}
	if len(this.m) >= this.capacity {
		e := this.freq[this.minFreq].Front()
		k := e.Value.(int)
		delete(this.m, k)
		delete(this.ptr, k)
		this.freq[this.minFreq].Remove(e)
	}
	this.m[key] = value{v, 1}
	if this.freq[1] == nil {
		this.freq[1] = list.New()
	}
	this.freq[1].PushBack(key)
	this.ptr[key] = this.freq[1].Back()
	this.minFreq = 1
}
