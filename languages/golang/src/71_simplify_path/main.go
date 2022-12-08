package _71_simplify_path

import "strings"

// must always begin with a slash /
// must be only a single slash / between two directory names
// last directory name (if it exists) must not end with a trailing /

func simplifyPath(path string) string {
	res := []string{}
	cursor := 0

	token, ok := getToken(path, &cursor)
	for ok {
		switch token {
		case ".":
		case "":
			// do nothing
		case "..":
			//  go back
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		default:
			res = append(res, token)
		}

		token, ok = getToken(path, &cursor)
	}
	// append
	resPath := strings.Join(res, "/")
	return "/" + resPath
}

func getToken(path string, cursor *int) (string, bool) {
	if *cursor >= len(path) {
		return "", false
	}
	// find next /
	var i int
	for i = *cursor; i < len(path); i++ {
		if path[i] == '/' {
			break
		}
	}
	token := path[*cursor:i]
	*cursor = i + 1

	return token, true
}
