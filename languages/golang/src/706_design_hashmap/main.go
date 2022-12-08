package _706_design_hashmap

const (
	maxCap = 1000000
	slotCt = 100
)

type MyHashMap struct {
	data [slotCt][]int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	ind := key % slotCt
	if this.data[ind] == nil {
		this.data[ind] = make([]int, maxCap/slotCt+1)
		for i := range this.data[ind] {
			this.data[ind][i] = -1
		}
	}
	this.data[ind][key/slotCt] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	ind := key % slotCt
	if this.data[ind] == nil {
		return -1
	}
	return this.data[ind][key/slotCt]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	ind := key % slotCt
	if this.data[ind] == nil {
		return
	}
	this.data[ind][key/slotCt] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
