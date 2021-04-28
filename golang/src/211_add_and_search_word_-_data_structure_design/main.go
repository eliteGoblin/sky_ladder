package _11_add_and_search_word___data_structure_design

type WordDictionary struct {
	has bool
	mp map[byte]*WordDictionary
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		mp: make(map[byte]*WordDictionary),
	}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
	cur := this
	for i := range word {
		ch := word[i]
		if _, has := cur.mp[ch]; !has {
			cur.mp[ch] = &WordDictionary{
				mp: make(map[byte]*WordDictionary),
			}
		}
		cur = cur.mp[ch]
	}
	cur.has = true
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	if word == "" {
		return this.has
	}
	if word[0] != '.' {
		if _, has := this.mp[word[0]]; !has {
			return false
		}
		return this.mp[word[0]].Search(word[1:])
	}
	for _, v := range this.mp {
		found := v.Search(word[1:])
		if found {
			return true
		}
	}
	return false
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */