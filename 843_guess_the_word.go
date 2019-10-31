package amazon

import "math/rand"

func findSecretWord(wordlist []string, master *Master) {
	lst := wordlist
	for len(lst) > 0 {
		word := lst[rand.Int()%len(lst)]
		diff := master.Guess(word)
		if diff == 6 {
			return
		}
		nextList := make([]string, 0, len(lst))
		for i := 0; i < len(lst); i++ {
			if distance(word, lst[i]) == diff {
				nextList = append(nextList, lst[i])
			}
		}
		lst = nextList
	}
}

func distance(a, b string) int {
	ct := 0
	for i := range a {
		if a[i] == b[i] {
			ct++
		}
	}
	return ct
}
