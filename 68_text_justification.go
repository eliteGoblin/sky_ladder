package amazon

// ["This", "is", "an", "example", "of", "text", "justification."]
// [
//    "This    is    an",
//    "example  of text",
//    "justification.  "
// ]
func fullJustify(words []string, maxWidth int) []string {
	res := make([]string, 0)
	for i := 0; i < len(words); {
		wordsOfLine := getWordsOfLine(words[i:], maxWidth)
		i += len(wordsOfLine)
		var line string
		if i == len(words) {
			line = getLastLine(wordsOfLine, maxWidth)
		} else {
			line = getLine(wordsOfLine, maxWidth)
		}
		res = append(res, line)
	}
	return res
}

func getWordsOfLine(words []string, maxWidth int) []string {
	res := make([]string, 0)
	for i := range words {
		if maxWidth <= 0 {
			break
		}
		wordLen := len(words[i])
		if i > 0 {
			wordLen++
		}
		if wordLen <= maxWidth {
			res = append(res, words[i])
		}
		maxWidth -= wordLen
	}
	return res
}

func getLine(words []string, maxWidth int) string {
	if len(words) == 1 {
		return getLastLine(words, maxWidth)
	}
	spaces := make([]string, len(words)-1)
	allSpaces := maxWidth
	for _, v := range words {
		allSpaces -= len(v)
	}
	for i := 0; allSpaces > 0; {
		spaces[i] += " "
		i = (i + 1) % len(spaces)
		allSpaces--
	}
	res := ""
	for i := 0; i < len(words)-1; i++ {
		res += words[i] + spaces[i]
	}
	res += words[len(words)-1]
	return res
}

func getLastLine(words []string, maxWidth int) string {
	res := words[0]
	for i := 1; i < len(words); i++ {
		res += " " + words[i]
	}
	maxWidth -= len(res)
	for maxWidth > 0 {
		res += " "
		maxWidth--
	}
	return res
}
