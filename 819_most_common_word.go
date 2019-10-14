package A

import "strings"

func mostCommonWord(paragraph string, banned []string) string {
	banDict := make(map[string]struct{})
	for _, v := range banned {
		banDict[v] = struct{}{}
	}
	ct := make(map[string]int)
	maxCt := 0
	res := ""
	for i := 0; i < len(paragraph); {
		next := strings.IndexAny(paragraph[i:], " !?',;.")
		if next == -1 {
			next = len(paragraph)
		} else {
			next += i
		}
		if next == i {
			i++
			continue
		}
		word := strings.ToLower(paragraph[i:next])
		if _, ok := banDict[word]; !ok {
			ct[word]++
			if ct[word] > maxCt {
				maxCt = ct[word]
				res = word
			}
		}
		if next < 0 {
			break
		}
		i = next + 1
	}
	return res
}
