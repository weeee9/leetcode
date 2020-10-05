package leetcode

import "strings"

func simplifyPath(path string) string {
	stack := []string{}
	arr := strings.Split(path, "/")
	for _, p := range arr {
        p = strings.Replace(p, " ", "", -1)
		if p == ".." {
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]    
            }			
		} else if p != "." && len(p) > 0 {
			stack = append(stack, p)
		}
	}

	if len(stack) == 0 {
		return "/"
	}

	return "/" + strings.Join(stack, "/")
}
