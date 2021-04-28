package _08_implement_trie

type Trie struct {
	has bool
	mp  map[byte]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		mp:  make(map[byte]*Trie),
		has: true,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for i := range word {
		ch := word[i]
		if _, has := cur.mp[ch]; !has {
			cur.mp[ch] = &Trie{
				mp: make(map[byte]*Trie),
			}
		}
		cur = cur.mp[ch]
	}
	cur.has = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if word == "" {
		return this.has
	}
	if _, has := this.mp[word[0]]; !has {
		return false
	}
	return this.mp[word[0]].Search(word[1:])
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return true
	}
	if _, has := this.mp[prefix[0]]; !has {
		return false
	}
	return this.mp[prefix[0]].StartsWith(prefix[1:])
}
