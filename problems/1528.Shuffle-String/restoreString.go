package leetcode

func restoreString(s string, indices []int) string {
	b := make([]byte, len(s))

	for i, pos := range indices {
		b[pos] = s[i]
	}

	return string(b)
}
