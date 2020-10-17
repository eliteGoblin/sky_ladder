package _211_design_add_and_search_words_data_structure

type TrieNode struct {
	isWord bool
	mp     map[rune]*TrieNode
}

type WordDictionary struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		root: &TrieNode{
			mp:     make(map[rune]*TrieNode),
			isWord: false,
		},
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	p := this.root
	for _, v := range word {
		if _, ok := p.mp[v]; !ok {
			p.mp[v] = &TrieNode{
				mp:     make(map[rune]*TrieNode),
				isWord: false,
			}
		}
		p = p.mp[v]
	}
	p.isWord = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return search(word, this.root)
}

func search(word string, root *TrieNode) bool {
	if word == "" {
		return root.isWord
	}
	if word[0] == '.' {
		for _, v := range root.mp {
			if search(word[1:], v) {
				return true
			}
		}
		return false
	}
	if _, ok := root.mp[rune(word[0])]; !ok {
		return false
	}
	return search(word[1:], root.mp[rune(word[0])])
}
