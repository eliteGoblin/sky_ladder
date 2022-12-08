package _767_reorganize_string

import (
	"fmt"
	"sort"
)

type node struct {
	char byte
	ct   int
}

type sorter []node

func (ps sorter) Len() int {
	return len(ps)
}
func (ps sorter) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
func (ps sorter) Less(i, j int) bool {
	return ps[i].ct > ps[j].ct
}

func reorganizeString(S string) string {
	arr := make([]node, 26)
	for i := range S {
		arr[S[i]-'a'].char = S[i]
		arr[S[i]-'a'].ct++
	}
	sort.Sort(sorter(arr))
	var last int
	fmt.Println(arr)
	for i := range arr {
		if arr[i].ct > (len(S)+1)/2 {
			return ""
		}
		if arr[i].ct == 0 {
			break
		}
		last = i
	}
	arr = arr[:last+1]
	return recur(arr, "")
}

func recur(arr []node, preResult string) string {
	if len(arr) == 0 {
		return preResult
	}
	if len(arr) == 1 {
		return preResult + string(arr[0].char)
	}
	for arr[0].ct > 0 && arr[1].ct > 0 {
		preResult += string(arr[0].char)
		preResult += string(arr[1].char)
		arr[0].ct--
		arr[1].ct--
	}
	nextArr := arr[2:]
	if arr[0].ct > 0 {
		nextArr = append(nextArr, arr[0])
	}
	if arr[1].ct > 0 {
		nextArr = append(nextArr, arr[1])
	}
	sort.Sort(sorter(nextArr))
	return recur(nextArr, preResult)
}
