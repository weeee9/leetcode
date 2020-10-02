package leetcode

import (
	"sort"
)

func longestWord(words []string) string {
	sort.Strings(words)
	record := make(map[string]bool)
	var result string
	for _, word := range words {
		size := len(word)
		if size == 1 || record[word[:size-1]] {
			if size > len(result) {
				result = word
			}
			record[word] = true
		}
	}
	return result
}
