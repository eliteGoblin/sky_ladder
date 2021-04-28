package _165_compare_version_numbers

import "strconv"

func compareVersion(version1 string, version2 string) int {
	var part1, part2 int64
	i, j := 0, 0
	for i < len(version1) || j < len(version2) {
		part1, i = getToken(version1, i)
		part2, j = getToken(version2, j)
		if part1 != part2 {
			if part1 > part2 {
				return 1
			}
			return -1
		}
	}
	return 0
}

func getToken(input string, startIndex int) (token int64, nextIndex int) {
	if startIndex >= len(input) {
		return 0, len(input)
	}
	var i int
	for i = startIndex; i < len(input); i++ {
		if input[i] == '.' {
			break
		}
	}
	token, _ = strconv.ParseInt(input[startIndex:i], 10, 64)
	return token, i + 1
}
