package amazon

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	res := make([][]string, 0)
	mp := make(map[string]bool)
	paths := make([][]string, 0)
	path := make([]string, 0)
	words := make(map[string]bool)

	path = append(path, beginWord)
	paths = append(paths, append([]string{}, path...))
	level, minLevel := 1, 9999999

	for _, v := range wordList {
		mp[v] = true
	}

	for len(paths) > 0 {
		t := paths[0]
		paths = paths[1:]

		if len(t) > level {
			for k := range words {
				delete(mp, k)
			}
			words = make(map[string]bool)
			level = len(t)
			if level > minLevel {
				break
			}
		}
		last := t[len(t)-1]
		for i := range last {
			newLast := last
			for j := 'a'; j <= 'z'; j++ {
				newLast = last[:i] + string(j) + last[i+1:]
				if _, ok := mp[newLast]; !ok {
					continue
				}
				words[newLast] = true
				newPath := append([]string{}, t...)
				newPath = append(newPath, newLast)
				if newLast == endWord {
					res = append(res, append([]string{}, newPath...))
					minLevel = level
				} else {
					paths = append(paths, append([]string{}, newPath...))
				}
			}
		}
	}
	return res
}
