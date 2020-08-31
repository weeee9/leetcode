package leetcode

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	open := map[rune]int{'(': 1, '{': 2, '[': 3}
	closing := map[rune]int{')': 1, '}': 2, ']': 3}

	stack := []rune{}

	for _, r := range s {
		if _, ok := open[r]; ok {
			stack = append(stack, r)
		} else if _, ok := closing[r]; ok {
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			if open[pop] != closing[r] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
