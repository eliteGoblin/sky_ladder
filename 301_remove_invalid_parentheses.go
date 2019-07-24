package amazon

import (
	"container/list"
)

func removeInvalidParentheses(s string) []string {
	res := make([]string, 0)
	visited := make(map[string]bool)
	visited[s] = true
	lst := list.New()
	lst.PushBack(s)
	found := false

	for lst.Len() > 0 {
		e := lst.Front()
		lst.Remove(e)
		cur := e.Value.(string)
		if isValidParentheses(cur) {
			res = append(res, cur)
			found = true
		}
		if found {
			continue
		}
		for i := range cur {
			if cur[i] != '(' && cur[i] != ')' {
				continue
			}
			eraseOne := cur[:i] + cur[i+1:] // 之前写成 eraseOne := s[:i] + s[i+1:], 极难发现!!!
			if _, ok := visited[eraseOne]; !ok {
				lst.PushBack(eraseOne)
				visited[eraseOne] = true
			}
		}
	}
	if len(res) == 0 {
		res = []string{""}
	}
	return res
}

func isValidParentheses(s string) bool {
	count := 0
	for i := range s {
		if s[i] == '(' {
			count++
		} else if s[i] == ')' {
			count--
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}

// my initial thought

// func removeInvalidParentheses(s string) []string {
// 	resMap := remove(s)
// 	res := make([]string, 0, len(resMap))
// 	for k := range resMap {
// 		res = append(res, k)
// 	}
// 	return res
// }

// func remove(s string) map[string]bool {
// 	res := make(map[string]bool)
// 	if check(s, 0, 0) {
// 		res[s] = true
// 		return res
// 	}
// 	for i := range s {
// 		if s[i] != '(' && s[i] != ')' {
// 			continue
// 		}
// 		curS := erase(s, i)
// 		if check(curS, 0, 0) {
// 			res[curS] = true
// 		}
// 	}
// 	if len(res) == 0 {
// 		for i := range s {
// 			res = merge(res, remove(erase(s, i)))
// 		}
// 	}
// 	return res
// }

// func merge(mp1, mp2 map[string]bool) map[string]bool {
// 	for k, v := range mp2 {
// 		mp1[k] = v
// 	}
// 	return mp1
// }

// func erase(s string, i int) string {
// 	return s[:i] + s[i+1:]
// }

// func check(s string, left, right int) bool {
// 	if right > left {
// 		return false
// 	}
// 	if s == "" {
// 		return left == right
// 	}
// 	if s[0] == '(' {
// 		return check(s[1:], left+1, right)
// 	} else if s[0] == ')' {
// 		return check(s[1:], left, right+1)
// 	}
// 	return check(s[1:], left, right)
// }
