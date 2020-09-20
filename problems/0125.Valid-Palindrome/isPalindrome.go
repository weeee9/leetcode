package leetcode

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "", -1)

	i, j := 0, len(s)-1
	for i < j {
		if !(97 <= s[i] && s[i] <= 122) && !(48 <= s[i] && s[i] <= 57) {
			i++
			continue
		}
		if !(97 <= s[j] && s[j] <= 122) && !(48 <= s[j] && s[j] <= 57) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
