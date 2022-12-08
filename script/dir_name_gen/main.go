package main

import (
	"fmt"
	"os"
	"strings"
)

func name(s string) string {
	s = strings.ToLower(s)
	bt := []byte(s)
	res := make([]byte, 0)
	for i := range bt {
		if bt[i] == '.' {
			continue
		}
		if strings.IndexByte(" (),-", bt[i]) != -1 {
			if len(res) > 0 && res[len(res)-1] == '_' {
				continue
			}
			res = append(res, '_')
		} else {
			res = append(res, bt[i])
		}
	}
	if res[len(res)-1] == '_' {
		res = res[:len(res)-1]
	}
	return string(res)
}

func main() {
	fmt.Println(name(os.Args[1]))
	return
}
