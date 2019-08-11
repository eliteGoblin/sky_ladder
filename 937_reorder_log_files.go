package amazon

import (
	"sort"
	"strings"
)

type alphaString []string

func (a alphaString) Len() int {
	return len(a)
}

func (a alphaString) Less(i, j int) bool {
	si := a[i][getFirstWordPos(a[i]):]
	sj := a[j][getFirstWordPos(a[j]):]
	if si == sj {
		return a[i][:getFirstWordPos(a[i])] < a[j][:getFirstWordPos(a[j])]
	}
	return si < sj
}

func (a alphaString) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func reorderLogFiles(logs []string) []string {
	alphaLogs := make([]string, 0)
	digiLogs := make([]string, 0)
	for _, v := range logs {
		pos := getFirstWordPos(v)
		if v[pos] >= 'a' && v[pos] <= 'z' {
			alphaLogs = append(alphaLogs, v)
		} else {
			digiLogs = append(digiLogs, v)
		}
	}
	res := make([]string, 0, len(logs))
	sort.Sort(alphaString(alphaLogs))
	res = append(res, alphaLogs...)
	res = append(res, digiLogs...)
	return res
}

func getFirstWordPos(log string) int {
	return strings.Index(log, " ") + 1
}
