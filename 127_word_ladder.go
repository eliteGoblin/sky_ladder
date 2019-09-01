package amazon

import "container/list"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	lst := list.New()
	lst.PushBack(beginWord)
	stepCount := make(map[string]int)
	stepCount[beginWord] = 1
	wordSet := make(map[string]bool)
	for _, v := range wordList {
		wordSet[v] = true
	}
	for lst.Len() > 0 {
		e := lst.Front()
		word := e.Value.(string)
		lst.Remove(e)
		if word == endWord {
			return stepCount[word]
		}
		for i := 0; i < len(word); i++ {
			for j := 'a'; j <= 'z'; j++ {
				newWord := word[:i] + string(j) + word[i+1:]
				_, ok1 := wordSet[newWord]
				_, ok2 := stepCount[newWord]
				if ok1 && !ok2 {
					lst.PushBack(newWord)
					stepCount[newWord] = stepCount[word] + 1
				}
			}
		}
	}
	return 0
}
