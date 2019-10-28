package amazon

import "fmt"

func compress(chars []byte) int {
	cur := 0
	for i, j := 0, 0; i < len(chars); i = j {
		for j < len(chars) && char[j] == char[i] {
			j++
		}
		chars[cur] = chars[i]
		cur++
		if j-i == 1 {
			continue
		}
		str := fmt.Sprintf("%d", j-i)
		for k := 0; k < len(str); k++ {
			chars[cur] = str[k]
			cur++
		}
	}
	return cur
}
