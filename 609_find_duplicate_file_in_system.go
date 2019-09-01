package amazon

import "strings"

func findDuplicate(paths []string) [][]string {
	mp := make(map[string][]string)
	for _, v := range paths {
		parse(v, mp)
	}
	res := make([][]string, 0)
	for _, v := range mp {
		if len(v) > 1 {
			res = append(res, v)
		}
	}
	return res
}

// root/a 1.txt(abcd) 2.txt(efgh)
func parse(s string, mp map[string][]string) {
	arr := strings.Split(s, " ")
	path := arr[0]
	for i := 1; i < len(arr); i++ {
		pos := strings.Index(arr[i], "(")
		name := arr[i][:pos]
		content := arr[i][pos+1 : len(arr[i])-1]
		if _, ok := mp[content]; ok {
			mp[content] = append(mp[content], path+"/"+name)
		} else {
			mp[content] = append([]string{}, path+"/"+name)
		}
	}
}
