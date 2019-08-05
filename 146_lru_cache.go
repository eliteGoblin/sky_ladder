package amazon

import "container/list"

type LRUCache struct {
	mp       map[int]*list.Element
	list     *list.List
	capacity int
}

func LRUConstructor(capacity int) LRUCache {
	return LRUCache{
		mp:       make(map[int]*list.Element),
		list:     list.New(),
		capacity: capacity,
	}
}

type element struct {
	key   int
	value int
}

func (this *LRUCache) Get(key int) int {
	if e, ok := this.mp[key]; ok {
		this.list.MoveToFront(e)
		return e.Value.(*element).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if e, ok := this.mp[key]; ok {
		this.list.MoveToFront(e)
		e.Value.(*element).value = value
		return
	}
	if this.list.Len() >= this.capacity {
		tail := this.list.Back()
		key := tail.Value.(*element).key
		delete(this.mp, key)
		this.list.Remove(tail)
	}
	this.list.PushFront(&element{
		key:   key,
		value: value,
	})
	this.mp[key] = this.list.Front()
}
