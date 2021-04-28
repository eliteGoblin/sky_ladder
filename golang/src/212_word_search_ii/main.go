package _12_word_search_ii

type TrieNode struct {
	child [26]*TrieNode
	str string
}

func NewTrieNode() *TrieNode{
	return &TrieNode{
	}
}

type Trie struct{
	root *TrieNode
}

func NewTrie() *Trie{
	return &Trie{
		root: NewTrieNode(),
	}
}

func (trie *Trie)Insert(str string) {
	p := trie.root
	for i := range str {
		index := str[i] - 'a'
		if p.child[index] == nil {
			p.child[index] = NewTrieNode()
		}
		p = p.child[index]
	}
	p.str = str
}

func findWords(board [][]byte, words []string) []string {
	res := make([]string, 0, 64)
	if len(board) == 0 || len(words) == 0 {
		return res
	}
	var visit [26][26]bool
	trie := NewTrie()
	for _, v := range words {
		trie.Insert(v)
	}
	for i := 0; i < len(board); i ++ {
		for j := 0; j < len(board[i]); j ++ {
			index := board[i][j] - 'a'
			if trie.root.child[index] != nil {
				search(board, trie.root.child[index], i, j, visit, &res)
			}
		}
	}
	return res
}

func search(board [][]byte, p *TrieNode, i, j int, visit [26][26]bool, res *[]string){

	if p.str != "" {
		*res = append(*res, p.str)
		p.str = ""
	}
	var dir = [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	visit[i][j] = true
	for _, d := range dir {
		ni := d[0] + i
		nj := d[1] + j
		if ni < 0 || ni >= len(board) || nj < 0 || nj >= len(board[0]) || visit[ni][nj]  {
			continue
		}
		index := board[ni][nj] - 'a'
		if p.child[index] != nil {
			search(board, p.child[index], ni, nj, visit, res)
		}

	}
	visit[i][j] = false
}