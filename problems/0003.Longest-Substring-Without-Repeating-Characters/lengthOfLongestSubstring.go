package leetcode

import "math"

func lengthOfLongestSubstring(s string) int {
	letters, head, tail := make(map[byte]int), 0, -1
	result := 0
	for head < len(s) {
		if _, ok := letters[s[head]]; ok {
			tail = int(math.Max(float64(tail), float64(letters[s[head]])))
		}
		letters[s[head]] = head
		result = int(math.Max(float64(result), float64(head-tail)))
		head++
	}
	return result
}
